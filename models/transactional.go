package models

type (

	// Transactional ...
	Transactional struct {
		Destination `json:"destination" validate:"required"`
		Message     `json:"message" validate:"required"`
	}

	// Message return object collection of body and subject for message
	Message struct {
		MessageBody    `json:"body" validate:"required"`
		MessageSubject `json:"subject" validate:"required"`
	}

	// Destination return objects of receiver collections addresses
	Destination struct {
		CcAddresses []string `json:"cc_addresses"`
		ToAddresses []string `json:"to_addresses" validate:"required"`
	}

	// MessageSubject return object of subject data in mail
	MessageSubject struct {
		Charset string `json:"charset" validate:"required"`
		Data    string `json:"data" validate:"required"`
	}

	// MessageBody return object of body data in mail
	MessageBody struct {
		MessageBodyHTML `json:"html"`
		MessageBodyText `json:"text"`
	}

	// MessageBodyHTML return object of message body in html format
	MessageBodyHTML struct {
		Charset string `json:"charset"`
		Data    string `json:"data"`
	}

	// MessageBodyText return object of message body in raw text
	MessageBodyText struct {
		Charset string `json:"charset"`
		Data    string `json:"data"`
	}
)
