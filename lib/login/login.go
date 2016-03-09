package login

import (
    "github.com/headzoo/surf"
    "errors"
)

//Tokens type
type Tokens struct {
    Xt string
    SID string
}


//Login type
type Login struct {
    Email string
    Passwd string
}

//GetSID to get the credentials for the user
func (l *Login) GetSID() (string, error) {
    bow := surf.NewBrowser()
    bow.AddRequestHeader("service", "sj")
    bow.AddRequestHeader("continue", "https://play.google.com/music/listen")

    err := bow.Open("https://accounts.google.com/ServiceLoginAuth")
    if err != nil {
        return "", errors.New("Error opening browser.")
    }
    
    fm, err := bow.Form("[id='gaia_loginform']")
    if err != nil {
        return "", errors.New("Error getting form.")
    }
    
    err = fm.Input("Email", l.Email)
    if err != nil {
        return "", err
    }
    
    err = fm.Submit()
    if err != nil {
        return "", errors.New("Error submitting form.")
    }
    
    fm, err = bow.Form("form")
    if err != nil {
        return "", err
    }
    
    err = fm.Input("Passwd", l.Passwd)
    if err != nil {
        return "", err
    }
    err = fm.Submit()
    if err != nil {
        return "", err
    }
    
    cookies := bow.SiteCookies()
    for _, v := range cookies {
        if v.Name == "SID" {
            return v.Value, nil
        }
    }
    
    return "", nil
}