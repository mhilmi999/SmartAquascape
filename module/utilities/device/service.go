package device

import "errors"

type Service interface {
	GetAllDevice() ([]Device, error)
	GetDeviceByID(ID int) (Device, error)
	GetAllTemperatures() ([]Temperature, error)
	GetAllWaterLevel() ([]WaterLevel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllDevice() ([]Device, error) {
	device, err := s.repository.GetAllDevice()
	if err != nil {
		return device, err
	}

	return device, nil
}

func (s *service) GetDeviceByID(ID int) (Device, error) {
	device, err := s.repository.FindByID(ID)
	if err != nil {
		return device, err
	}

	if device.No_device == 0 {
		return device, errors.New("No Device found on with that ID")
	}

	return device, nil
}

func (s *service) GetAllTemperature() ([]Temperature, error) {
	temperature, err := s.repository.GetAllTemperatures()
	if err != nil {
		return temperature, err
	}

	return temperature, nil
}

func (s *service) GetAllWaterLevel() ([]WaterLevel, error) {
	waterLevel, err := s.repository.GetAllWaterLevel()
	if err != nil {
		return waterLevel, err
	}

	return waterLevel, nil
}
