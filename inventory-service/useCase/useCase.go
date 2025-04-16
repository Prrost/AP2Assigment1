package useCase

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"inventory-service/Storage"
	"inventory-service/config"
	"inventory-service/domain"
	"log"
)

type UseCase struct {
	storage Storage.Storage
	cfg     *config.Config
}

func NewUseCase(storage Storage.Storage, cfg *config.Config) *UseCase {
	return &UseCase{
		storage: storage,
		cfg:     cfg,
	}
}

func (u *UseCase) CreateProduct(toInsertObject domain.Object) (domain.Object, error) {
	const op = "CreateProduct"

	var object domain.Object

	//validation of object
	switch {
	case toInsertObject.Name == "":
		log.Printf("[%s] Name is empty", op)
		return domain.Object{}, status.Error(codes.InvalidArgument, "Name is empty")
	case toInsertObject.Amount < -1:
		log.Printf("[%s] Amount is negative", op)
		return domain.Object{}, status.Error(codes.InvalidArgument, "Amount is negative")
	case toInsertObject.Amount == 0:
		log.Printf("[%s] Amount is zero", op)
		return domain.Object{}, status.Error(codes.InvalidArgument, "Amount is zero")
	}

	//amount check
	if toInsertObject.Amount == -1 {
		object.Amount = 0
		object.Available = false
	} else {
		object.Amount = toInsertObject.Amount
		object.Available = true
	}

	object.Name = toInsertObject.Name

	//db creation
	outObject, err := u.storage.CreateObject(object)
	if err != nil {
		if errors.Is(err, Storage.ErrAlreadyExists) {
			log.Printf("[%s] Object already exists", op)
			return domain.Object{}, status.Error(codes.AlreadyExists, "Object already exists")
		}
		log.Printf("[CreateProduct] failed to create product: %v", err)
		return domain.Object{}, status.Error(codes.Internal, err.Error())
	}

	return outObject, nil
}

func (u *UseCase) GetProductById(id int64) (domain.Object, error) {
	const op = "GetProductById"

	if id <= 0 {
		log.Printf("[%s] ID is not valid", op)
		return domain.Object{}, status.Error(codes.InvalidArgument, "ID is not valid")
	}

	object, err := u.storage.GetObjectByID(int(id))
	if err != nil {
		if errors.Is(err, Storage.ErrNotFound) {
			log.Printf("[%s] Object not found", op)
			return domain.Object{}, status.Error(codes.NotFound, "Object not found")
		}
		log.Printf("[%s] failed to get object by id: %v", op, err)
		return domain.Object{}, status.Error(codes.Internal, err.Error())
	}

	return object, nil
}

func (u *UseCase) GetAllProducts(name string, limit int32, offset int32) ([]domain.Object, error) {
	const op = "GetAllProducts"

	if limit <= 0 {
		limit = 5
	}
	if offset < 0 {
		offset = 0
	}

	objects, err := u.storage.ListProducts(name, int(limit), int(offset))
	if err != nil {
		log.Printf("[%s] failed to list products: %v", op, err)
		return []domain.Object{}, status.Error(codes.Internal, err.Error())
	}

	return objects, nil
}

func (u *UseCase) UpdateProductById(toUpdateObject domain.Object) (domain.Object, error) {
	const op = "UpdateProductById"

	//isExists
	object, err := u.storage.GetObjectByID(toUpdateObject.ID)
	if err != nil {
		if errors.Is(err, Storage.ErrNotFound) {
			log.Printf("[%s] Object to update not found", op)
			return domain.Object{}, status.Error(codes.NotFound, "Object not found")
		}
		log.Printf("[%s] failed to get object by id: %v", op, err)
		return domain.Object{}, status.Error(codes.Internal, err.Error())
	}

	//name collision
	if toUpdateObject.Name != object.Name {
		ok, err := u.storage.IsProductExists(toUpdateObject.Name)
		if err != nil {
			log.Printf("[%s] failed to check if object exists: %v", op, err)
			return domain.Object{}, status.Error(codes.Internal, err.Error())
		}
		if ok {
			log.Printf("[%s] object already exists, not updating", op)
			return domain.Object{}, status.Error(codes.AlreadyExists, "Object with this name is already exists")
		}
	}

	//validation
	switch {
	case toUpdateObject.Amount < -1:
		log.Printf("[%s] Amount is negative", op)
		return domain.Object{}, status.Error(codes.InvalidArgument, "Amount is negative")
	case toUpdateObject.Amount == -1:
		object.Amount = 0
		object.Available = false
	case toUpdateObject.Amount > 0:
		object.Amount = toUpdateObject.Amount
		object.Available = true
	}

	if toUpdateObject.Name != "" {
		object.Name = toUpdateObject.Name
	}

	//updating db
	newObject, err := u.storage.UpdateObjectByID(toUpdateObject.ID, object)
	if err != nil {
		log.Printf("[%s] failed to update object: %v", op, err)
		return domain.Object{}, status.Error(codes.Internal, err.Error())
	}

	return newObject, nil
}

func (u *UseCase) DeleteProductById(id int64) (domain.Object, error) {
	const op = "DeleteProductById"

	if id <= 0 {
		log.Printf("[%s] ID is not valid", op)
		return domain.Object{}, status.Error(codes.InvalidArgument, "ID is not valid")
	}

	object, err := u.storage.DeleteObjectByID(int(id))
	if err != nil {
		if errors.Is(err, Storage.ErrNotFound) {
			log.Printf("[%s] Object not found", op)
			return domain.Object{}, status.Error(codes.NotFound, "Object not found")
		}
		log.Printf("[%s] failed to get object by id: %v", op, err)
		return domain.Object{}, status.Error(codes.Internal, err.Error())
	}

	return object, nil
}
