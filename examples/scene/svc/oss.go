package svc

// oss defined defined
type Oss interface {
	PutObjectFromFile(target string, src string) error
	GetObjectToFile(target string, src string) error
	ListObjects() ([]string, error)
	DeleteObject(string) error
}

// PutObjectFromFile defined TODO
func (svc *SvcHepler) PutObjectFromFile(target string, src string) error {
	err := svc.bucket.PutObjectFromFile(target, src)
	return err
}

// PutObjectFromFile defined TODO
func (svc *SvcHepler) GetObjectToFile(target string, src string) error {
	err := svc.bucket.GetObjectToFile(target, src)
	return err
}

// PutObjectFromFile defined TODO
func (svc *SvcHepler) ListObjects() (keys []string, err error) {
	lsRes, err := svc.bucket.ListObjects()
	if err != nil {
		return []string{}, err
	}
	for _, object := range lsRes.Objects {
		keys = append(keys, object.Key)
	}
	return keys, err
}

// PutObjectFromFile defined TODO
func (svc *SvcHepler) DeleteObject(target string) error {
	err := svc.bucket.DeleteObject(target)
	return err
}
