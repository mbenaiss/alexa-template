package main

type IntentType = string

// Intent is the type of the intent.
const (
	IntentTypeLaunch IntentType = "LaunchRequest"
)

// IntentName is the name of the intent.
type IntentName = string

// Standard Intent Names.
const (
	CancelIntent           IntentName = "AMAZON.CancelIntent"
	FallbackIntent         IntentName = "AMAZON.FallbackIntent"
	HelpIntent             IntentName = "AMAZON.HelpIntent"
	LoopOffIntent          IntentName = "AMAZON.LoopOffIntent"
	LoopOnIntent           IntentName = "AMAZON.LoopOnIntent"
	MoreIntent             IntentName = "AMAZON.MoreIntent"
	NavigateHomeIntent     IntentName = "AMAZON.NavigateHomeIntent"
	NavigateSettingsIntent IntentName = "AMAZON.NavigateSettingsIntent"
	NextIntent             IntentName = "AMAZON.NextIntent"
	NoIntent               IntentName = "AMAZON.NoIntent"
	PageDownIntent         IntentName = "AMAZON.PageDownIntent"
	PageUpIntent           IntentName = "AMAZON.PageUpIntent"
	PauseIntent            IntentName = "AMAZON.PauseIntent"
	PreviousIntent         IntentName = "AMAZON.PreviousIntent"
	RepeatIntent           IntentName = "AMAZON.RepeatIntent"
	ResumeIntent           IntentName = "AMAZON.ResumeIntent"
	ScrollDownIntent       IntentName = "AMAZON.ScrollDownIntent"
	ScrollLeftIntent       IntentName = "AMAZON.ScrollLeftIntent"
	ScrollRightIntent      IntentName = "AMAZON.ScrollRightIntent"
	ScrollUpIntent         IntentName = "AMAZON.ScrollUpIntent"
	SelectIntent           IntentName = "AMAZON.SelectIntent"
	SendToPhoneIntent      IntentName = "AMAZON.SendToPhoneIntent"
	ShuffleOffIntent       IntentName = "AMAZON.ShuffleOffIntent"
	ShuffleOnIntent        IntentName = "AMAZON.ShuffleOnIntent"
	StartOverIntent        IntentName = "AMAZON.StartOverIntent"
	StopIntent             IntentName = "AMAZON.StopIntent"
	YesIntent              IntentName = "AMAZON.YesIntent"
)

// AlexaRequest is the request from Alexa.
type AlexaRequest struct {
	Version string  `json:"version"`
	Request Request `json:"request"`
}

// Request is the request from Alexa.
type Request struct {
	Type   string `json:"type"`
	Time   string `json:"timestamp"`
	Intent Intent `json:"intent"`
}

// Intent is the request from Alexa.
type Intent struct {
	Name               string `json:"name"`
	ConfirmationStatus string `json:"confirmationStatus"`
	Slots              Slots  `json:"slots"`
}

// Slots is the request from Alexa.
type Slots struct {
	SlotName SlotName `json:"SlotName"`
}

// SlotName is the request from Alexa.
type SlotName struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
}

// AlexaResponse is the response to Alexa.
type AlexaResponse struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

// Response is the response to Alexa.
type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card        `json:"card,omitempty"`
	Reprompt         *Reprompt    `json:"reprompt,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

// Card is the object containing a card to render to the Amazon Alexa App.
type Card struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Image *Image `json:"image,omitempty"`
}

// Image is the response to Alexa.
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// OutputSpeech is the object containing the speech to render to the user.
type OutputSpeech struct {
	Type         string `json:"type,omitempty"`
	Text         string `json:"text,omitempty"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

// Reprompt is the object containing the outputSpeech to use if a re-prompt is necessary.
type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`
}

// SimpleResponse is the simple response to Alexa.
func SimpleResponse(text string, endSession bool) AlexaResponse {
	return AlexaResponse{
		Version: "1.0",
		Response: Response{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
			ShouldEndSession: endSession,
		},
	}
}

// ResponseWithCard is the response to Alexa.
func ResponseWithCard(title, text string, urlImage *string) AlexaResponse {
	resp := AlexaResponse{
		Version: "1.0",
		Response: Response{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
			Card: &Card{
				Type:  "Simple",
				Title: title,
				Text:  text,
			},
			ShouldEndSession: true,
		},
	}

	if urlImage != nil {
		resp.Response.Card.Image = &Image{
			SmallImageURL: *urlImage,
			LargeImageURL: *urlImage,
		}
	}

	return resp
}

// GetIntentName returns the intent name.
func GetIntentName(req AlexaRequest) IntentName {
	return IntentName(req.Request.Intent.Name)
}

// GetIntentType returns the intent type.
func GetIntentType(req AlexaRequest) IntentType {
	return IntentType(req.Request.Type)
}
