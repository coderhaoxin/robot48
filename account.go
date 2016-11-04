package main

import "github.com/parnurzeal/gorequest"
import "fmt"

func queryProfile() {
	res, body, err := gorequest.
		New().
		Set("User-Agent", userAgent).
		Set("Origin", "http://user.snh48.com").
		Set("Referer", "http://user.snh48.com/").
		Set("Cookie", userCookie).
		Post("http://user.snh48.com/default.php").
		Query("example=index").
		Send("act=default").
		End()

	if err != nil {
		fmt.Errorf("query profile error: %v", err)
	}
	fmt.Println(res, "\n\n", body, "\n\n")
	// fmt.Printf("profile: %s", body)
}

func getOrders() {
	res, body, err := gorequest.
		New().
		Set("User-Agent", userAgent).
		Set("Referer", "http://user.snh48.com/"). //
		Set("Cookie", cookie).
		Get("http://shop.snh48.com/user.php").
		Query("act=order_list").
		End()

	if err != nil {
		fmt.Errorf("get order list error: %v", err)
	}

	fmt.Println(res, "\n\n", body, "\n\n", cookie)
	// fmt.Printf("profile: %s", body)
}
