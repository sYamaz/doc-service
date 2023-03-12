package usecase

type (
	DocRepository interface {
		GetOwnDocs(userId string) ([]DocHeader, error)
		GetOwnDoc(userId string, documentId string) (*Doc, error)
		AppendOwnDoc(userId string, doc Doc) (documentId string, err error)
		UpdateOwnDoc(userId string, doc Doc) error
		DeleteOwnDoc(userId string, docId string) error
	}

	DocGetOutputPort interface {
		Failure(err error) error
		Success(doc *Doc) error
	}

	DocPostOutputPort interface {
		Failure(err error) error
		Success(docId string) error
	}

	DocListOutputPort interface {
		Failure(err error) error
		Success(docs []DocHeader) error
	}

	DocPutOutputPort interface {
		Failure(err error) error
		Success() error
	}

	DocDeleteOutputPort interface {
		Failure(err error) error
		Success() error
	}

	DocListCondition struct {
		OnlyOwn bool
	}

	DocService interface {
		GetDocs(userId string, condition *DocListCondition, out DocListOutputPort) error
		GetDoc(userId string, docId string, out DocGetOutputPort) error
		PostDoc(userId string, doc Doc, out DocPostOutputPort) error
		UpdateDoc(userId string, doc Doc, out DocPutOutputPort) error
		DeleteDoc(userId string, docId string, out DocDeleteOutputPort) error
	}

	docService struct {
		rep DocRepository
	}

	DocHeader struct {
		Id    string
		Title string
	}

	Doc struct {
		Id    string
		Title string
		Body  string
	}
)

func NewDocService(rep DocRepository) DocService {
	return &docService{rep: rep}
}

func (s *docService) GetDocs(userId string, condition *DocListCondition, out DocListOutputPort) error {
	docs, err := s.rep.GetOwnDocs(userId)
	if err != nil {
		out.Failure(err)
	}

	return out.Success(docs)
}

func (s *docService) GetDoc(userId string, docId string, out DocGetOutputPort) error {
	doc, err := s.rep.GetOwnDoc(userId, docId)
	if err != nil {
		out.Failure(err)
	}
	return out.Success(doc)
}

func (s *docService) PostDoc(userId string, doc Doc, out DocPostOutputPort) error {
	docId, err := s.rep.AppendOwnDoc(userId, doc)
	if err != nil {
		out.Failure(err)
	}
	return out.Success(docId)
}

func (s *docService) UpdateDoc(userId string, doc Doc, out DocPutOutputPort) error {
	err := s.rep.UpdateOwnDoc(userId, doc)
	if err != nil {
		out.Failure(err)
	}
	return out.Success()
}

func (s *docService) DeleteDoc(userId string, docId string, out DocDeleteOutputPort) error {
	err := s.rep.DeleteOwnDoc(userId, docId)
	if err != nil {
		out.Failure(err)
	}
	return out.Success()
}
