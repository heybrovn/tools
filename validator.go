package tools

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	vietnamese "github.com/go-playground/locales/vi"
	"github.com/heybrovn/tools/validator"
	vi_translations "github.com/heybrovn/tools/validator/translations/vi"
	"github.com/heybrovn/tools/validator/ut"
)

var validatetag = []string{
	"required",
	"email",
	"isStringLAlphabetic",  // kiểu dữ liệu string alpha
	"isStringAlphabetic",   //a-zA-Z0-9_
	"isStringAlphabeticV2", //a-zA-Z0-9_.\-
	"isStringNumber",       // 0-9
	"isStringName",         //[ 0-9a-zA-Z_áàảãạăắặằẳẵâấầẩẫậÁÀẢÃẠĂẮẶẰẲẴÂẤẦẨẪẬđĐéèẻẽẹêếềểễệÉÈẺẼẸÊẾỀỂỄỆíìỉĩịÍÌỈĨỊóòỏõọôốồổỗộơớờởỡợÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢúùủũụưứừửữựÚÙỦŨỤƯỨỪỬỮỰýỳỷỹỵÝỲỶỸỴ]
	"isPhoneNumber",
	"isVNPhoneNumber", // kiểu số điện thoại 84xxx
	"isUrl",
	"json",
	"url",
	"numeric", // là kiểu số
	"hexadecimal",
	"uppercase",
	"lowercase",
	"len",
	"min",
	"max",
	"eq",  // Equals
	"ne",  //Not Equal
	"gt",  //Greater Than
	"gte", //Greater Than or Equal
	"lt",  //Less Than
	"lte", //Less Than or Equal
}

var validatetagtext = map[string]string{
	"required":             "không được để trống",
	"email":                "không đúng định dạng",
	"isStringLAlphabetic":  "phải là ký tự thường",           // kiểu dữ liệu string alpha
	"isStringAlphabetic":   "phải là ký tự chữ hoặc số",      //a-zA-Z0-9_
	"isStringAlphabeticV2": "phải là ký tự chữ hoặc số",      //a-zA-Z0-9_.\-
	"isStringNumber":       "phải là số",                     // 0-9
	"isStringName":         "không được chứa ký tự đặc biệt", //[ 0-9a-zA-Z_áàảãạăắặằẳẵâấầẩẫậÁÀẢÃẠĂẮẶẰẲẴÂẤẦẨẪẬđĐéèẻẽẹêếềểễệÉÈẺẼẸÊẾỀỂỄỆíìỉĩịÍÌỈĨỊóòỏõọôốồổỗộơớờởỡợÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢúùủũụưứừửữựÚÙỦŨỤƯỨỪỬỮỰýỳỷỹỵÝỲỶỸỴ]
	"isPhoneNumber":        "không phải là số điện thoại",
	"isVNPhoneNumber":      "không phải là số điện thoại", // kiểu số điện thoại 84xxx
	"isUrl":                "không phải là địa chỉ website",
}

type Translates interface {
	Translates() map[string]string
}

type Messages interface {
	Messages() map[string]string
}
type NoValidate interface {
	NoValidate() map[string]bool
}

type StrucInfo struct {
	Name       string
	Translates map[string]string
	Messages   map[string]string
	NoValidate map[string]bool
}

func getStructInfo(s interface{}) StrucInfo {
	fieldMessages := map[string]string{}
	fieldTrans := map[string]string{}
	fieldNoValidate := map[string]bool{}

	fooType := reflect.TypeOf(s)

	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		switch method.Name {
		case "Translates":
			fieldTrans = s.(Translates).Translates()
		case "Messages":
			fieldMessages = s.(Messages).Messages()
		case "NoValidate":
			fieldNoValidate = s.(NoValidate).NoValidate()
		default:
			fmt.Println(method.Name)
		}
	}

	return StrucInfo{
		Name:       reflect.TypeOf(s).Name(),
		Translates: fieldTrans,
		Messages:   fieldMessages,
		NoValidate: fieldNoValidate,
	}

}

func isPhoneNumber(fl validator.FieldLevel) bool {

	if len(fl.Field().String()) > 11 || len(fl.Field().String()) <= 6 {
		return false
	}
	return true
}

// isVNPhoneNumber quy dịnh số điện thoại có đầu số 84xxxx
func isVNPhoneNumber(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) == 11
}

// lower alpabetic
func isStringLAlphabetic(fl validator.FieldLevel) bool {
	var checkRegexp = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	return checkRegexp(fl.Field().String())
}

func isStringAlphabetic(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) > 0 {
		var checkRegexp = regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`).MatchString
		return checkRegexp(fl.Field().String())
	}
	return true

}

func isStringAlphabeticV2(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) > 0 {
		var checkRegexp = regexp.MustCompile(`^[a-zA-Z0-9_.\- ]+$`).MatchString
		return checkRegexp(fl.Field().String())
	}
	return true

}

func isStringNumber(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) > 0 {
		var checkRegexp = regexp.MustCompile(`^[0-9]+$`).MatchString
		return checkRegexp(fl.Field().String())
	}
	return true
}

func isStringName(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) > 0 {
		var checkRegexp = regexp.MustCompile(`^[-@.;:,'\/#&+\w\s0-9a-zA-Z_áàảãạăắặằẳẵâấầẩẫậÁÀẢÃẠĂẮẶẰẲẴÂẤẦẨẪẬđĐéèẻẽẹêếềểễệÉÈẺẼẸÊẾỀỂỄỆíìỉĩịÍÌỈĨỊóòỏõọôốồổỗộơớờởỡợÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢúùủũụưứừửữựÚÙỦŨỤƯỨỪỬỮỰýỳỷỹỵÝỲỶỸỴ]+$`).MatchString
		return checkRegexp(fl.Field().String())
	}
	return true
}

func isUrl(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) > 0 {
		var checkRegexp = regexp.MustCompile(`^(http|https)(:\/\/)(www.)?(([a-zA-Z0-9\-]+\.){1,2}([a-zA-Z]{1,5}\.)?[a-zA-Z]{1,5})+$`).MatchString
		return checkRegexp(fl.Field().String())
	}
	return true
}

type NPvalidator struct {
	npvalidate *validator.Validate
	trans      ut.Translator
}

func NewNPvalidator() *NPvalidator {

	vie := vietnamese.New()
	uni := ut.New(vie, vie)
	tran, _ := uni.GetTranslator("vi")
	trans := tran
	npvalidate := validator.New()
	vi_translations.RegisterDefaultTranslations(npvalidate, tran)
	npV := &NPvalidator{
		npvalidate: npvalidate,
		trans:      trans,
	}
	npV.registerValidation()
	return npV

}

func (np *NPvalidator) registerValidation() {
	np.npvalidate.RegisterValidation("isPhoneNumber", isPhoneNumber)
	np.npvalidate.RegisterValidation("isVNPhoneNumber", isVNPhoneNumber)
	np.npvalidate.RegisterValidation("isStringLAlphabetic", isStringLAlphabetic)
	np.npvalidate.RegisterValidation("isStringAlphabetic", isStringAlphabetic)
	np.npvalidate.RegisterValidation("isStringAlphabeticV2", isStringAlphabeticV2)
	np.npvalidate.RegisterValidation("isStringNumber", isStringNumber)
	np.npvalidate.RegisterValidation("isStringName", isStringName)
	np.npvalidate.RegisterValidation("isUrl", isUrl)
}

func (np *NPvalidator) translateOverride(structinfo StrucInfo) {
	fncTrs := func(ut ut.Translator, fe validator.FieldError) string {
		fieldMessages := structinfo.Messages
		fieldTrans := structinfo.Translates
		field := fe.Field()
		vl, ok := fieldTrans[field]
		if !ok {
			vl = field
		}
		tag := fe.ActualTag()
		vltagMessages, errget := fieldMessages[field+"."+tag]
		if !errget {
			vltagMessages, errget = fieldMessages[tag]
			if !errget {
				vltagMessages, errget = validatetagtext[tag]
				if !errget {
					vltagMessages = "Lỗi chưa được định nghĩa " + tag
				}
			}
		}
		//vltagMessages = strings.ReplaceAll(vltagMessages, "{value}", fe.Value().(string))
		t, _ := ut.T(tag, vl, vltagMessages)
		return t
	}
	for _, strtag := range validatetag {
		fncRq := func(ut ut.Translator) error {
			// How to Fix error fatal error: concurrent map writes
			err := ut.Add(strtag, "{0} {1}", true) // // see universal-translator for details
			if err != nil {
				Logger.Infof("translateOverride add:", err, strtag, structinfo)
			}
			//panic("TEST")
			return err
		}
		//How to Fix error fatal error: concurrent map writes
		errR := np.npvalidate.RegisterTranslation(strtag, np.trans, fncRq, fncTrs)
		if errR != nil {
			Logger.Infof("RegisterTranslation:", errR, strtag)
		}
	}
}
func (np NPvalidator) ValidStruct(s interface{}) error {
	structInfo := getStructInfo(s)
	np.translateOverride(structInfo)
	err := np.npvalidate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fieldNoValidate := structInfo.NoValidate
			field := e.Field()
			tag := e.ActualTag()
			vl, ok := fieldNoValidate[field+"."+tag]
			if ok && vl {
				continue
			} else if !ok {
				vl, ok := fieldNoValidate[field]
				if ok && vl {
					continue
				}
			}
			return errors.New(e.Translate(np.trans))
		}
	}
	return nil
}
