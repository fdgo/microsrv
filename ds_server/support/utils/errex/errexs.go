package errex

import (
	"errors"
)
//--------------------------------------------------------
const (
	S_EMPTY_CODE = "Empty code."

	S_INVALID_ANSWER     = "Invalid answer."
	S_INVALID_CLASS      = "Invalid class."
	S_INVALID_CLASS_ID   = "Invalid class ID."
	S_INVALID_COURSEWARE = "Invalid courseware ID."
	S_INVALID_COVER      = "Invalid cover."
	S_INVALID_DATA       = "Invalid data."
	S_INVALID_GROUP      = "Invalid group."
	S_INVALID_ID         = "Invalid ID."
	S_INVALID_ISSUE      = "Invalid issue."
	S_INVALID_KEEPER     = "Invalid keeper."
	S_INVALID_KEY        = "Invalid key."
	S_INVALID_MEETING    = "Invalid meeting."
	S_INVALID_NAME       = "Invalid name."
	S_INVALID_NICKNAME   = "Invalid nickname."
	S_INVALID_PASSWORD   = "Invalid password."
	S_INVALID_PHONE      = "Invalid phone."
	S_INVALID_QUESTION   = "Invalid question."
	S_INVALID_RESPONSE   = "Invalid response."
	S_INVALID_SESSION    = "Invalid session."
	S_INVALID_STUDENT    = "Invalid student."
	S_INVALID_SUB_KEY    = "Invalid sub-key."
	S_INVALID_TOKEN      = "Invalid token."
	S_INVALID_TYPE       = "Invalid type."
	S_INVALID_USER       = "Invalid user."
	S_INVALID_VIDEO      = "Invalid video."

	S_NO_AUTHORITY  = "No authority."
	S_NO_COURSEWARE = "The courseware does not exist."
	S_NO_CLASS      = "The class does not exist."
	S_NO_EXAM       = "The exam does not exist."
	S_NO_GROUP      = "The group does not exist."
	S_NO_MEETING    = "The meeting does not exist."
	S_NO_USER       = "The user does not exist."
	S_NO_SERVICE    = "Service is unavailable."
	S_NO_VIDEO      = "The video does not exist."
)

//--------------------------------------------------------

var ERR_NO_DATABASE error = errors.New("No database available.")
var ERR_NO_CACHE error = errors.New("No cache available.")
var ERR_NO_QUEUE error = errors.New("No message queue available.")

var ERR_NO_AUTHORITY error = errors.New(S_NO_AUTHORITY)
var ERR_NO_COURSEWARE error = errors.New(S_NO_COURSEWARE)
var ERR_NO_CLASS error = errors.New(S_NO_CLASS)
var ERR_NO_EXAM error = errors.New(S_NO_EXAM)
var ERR_NO_GROUP error = errors.New(S_NO_GROUP)
var ERR_NO_MEETING error = errors.New(S_NO_MEETING)
var ERR_NO_SERVICE error = errors.New(S_NO_SERVICE)
var ERR_NO_USER error = errors.New(S_NO_USER)
var ERR_NO_VIDEO error = errors.New(S_NO_VIDEO)

var ERR_INVALID_CLASS error = errors.New(S_INVALID_CLASS)
var ERR_INVALID_GROUP error = errors.New(S_INVALID_GROUP)
var ERR_INVALID_ID error = errors.New(S_INVALID_ID)
var ERR_INVALID_MEETING error = errors.New(S_INVALID_MEETING)
var ERR_INVALID_NAME error = errors.New(S_INVALID_NAME)
var ERR_INVALID_NICKNAME error = errors.New(S_INVALID_NICKNAME)
var ERR_INVALID_PHONE error = errors.New(S_INVALID_PHONE)
var ERR_INVALID_RESPONSE error = errors.New(S_INVALID_RESPONSE)
var ERR_INVALID_SESSION error = errors.New(S_INVALID_SESSION)
var ERR_INVALID_TOKEN error = errors.New(S_INVALID_TOKEN)
var ERR_INVALID_TYPE error = errors.New(S_INVALID_TYPE)
var ERR_INVALID_USER error = errors.New(S_INVALID_USER)
var ERR_INVALID_VIDEO error = errors.New(S_INVALID_VIDEO)

//--------------------------------------------------------

var ERR_OUT_OF_TIME error = errors.New("Out of time.")

var ERR_MEETING_CLOSED error = errors.New("The meeting had been closed.")
var ERR_CLASS_CLOSED error = errors.New("The class had been closed.")

var ERR_CLASS_IS_NOT_EMPTY error = errors.New("The class is not empty.")
var ERR_GROUP_IS_NOT_EMPTY error = errors.New("The group is not empty.")
var ERR_CATEGORY_IS_NOT_EMPTY error = errors.New("The category is not empty.")

var ERR_DUPLICATED_ANSWER error = errors.New("Duplicated answer.")
var ERR_DUPLICATED_COURSEWARE error = errors.New("Duplicated courseware.")
var ERR_DUPLICATED_CLASS error = errors.New("Duplicated class.")
var ERR_DUPLICATED_GAODUN_COURSE_ID error = errors.New("Duplicated Gd course ID.")
var ERR_DUPLICATED_GROUP error = errors.New("Duplicated group.")
var ERR_DUPLICATED_MEETING error = errors.New("Duplicated meeting.")
var ERR_DUPLICATED_NAME error = errors.New("Duplicated name.")
var ERR_DUPLICATED_QUESTION error = errors.New("Duplicated question.")
var ERR_DUPLICATED_SCORE error = errors.New("Duplicated scores.")
var ERR_DUPLICATED_USER error = errors.New("Duplicated user.")
var ERR_DUPLICATED_VIDEO error = errors.New("Duplicated video.")

//--------------------------------------------------------

var ERR_NO_QUESTION error = errors.New("The question does not exist.")
var ERR_NO_FILE error = errors.New("The file does not exist.")
var ERR_NO_USER_ANSWER error = errors.New("The user answer does not exist.")
var ERR_NO_RECORD error = errors.New("The record does not exist.")

var ERR_INVALID_COURSEWARE error = errors.New("Invalid courseware.")
var ERR_INVALID_KEY error = errors.New("Invalid key.")
var ERR_INVALID_SOURCE error = errors.New("Invalid source.")
var ERR_INVALID_ANSWER error = errors.New("Invalid exam answer.")
var ERR_INVALID_SCORE error = errors.New("Invalid score.")
var ERR_INVALID_NOTE error = errors.New("Invalid note.")

var ERR_INVALID_QUESTION_TYPE error = errors.New("Invalid question type.")

//--------------------------------------------------------
