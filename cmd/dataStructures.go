// Copyright © 2017 <mattjones@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

// Sample data structure
//{
//    "name": {
//        "first": "",
//        "last": "",
//        "middle": "",
//        "nickname": "",
//        "alias": ""
//    },
//    "address": [{
//        "street 1": "",
//	  "street 2": "",
//	  "city": "",
//	  "state": "",
//	  "zipcode": "",
//	  "country": ""
//    }],
//    "email": {
//	  "secure": [""],
//	  "personal": [""],
//	  "work": [""],
//	  "other": [""]
//    },
//    "phone": {
//	  "mobile": [""],
//	  "home": [""],
//	  "work": [""],
//	  "other": [""]
//    },
//    "networks": {
//	  "bitbucket": "",
//	  "blog": "",
//	  "facebook": "",
//	  "github": "",
//	  "gitlab": "",
//	  "linkedin": "",
//	  "other": "",
//	  "site": "",
//	  "twitter": ""
//    },
//    "encryption": {
//	  "keys": [{
//	      "public key": "",
//	      "fingerprint": ""
//	  }]
//    },
//    "notes": [""]
//}

type Name struct {
	First     string `json:"first"`
	Last      string `json:"last"`
	Middle    string `json:"middle"`
	Nickname  string `json:"nick"`
	Alias     string `json:"alias"`

}
type Address struct {
	Street1  string `json:"s1"`
	Street2  string `json:"s2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zip"`
	Country  string `json:"country"`
}

type Email struct {
	Secure    []string `json:"secure"`
	Personal  []string `json:"personal"`
	Work      []string `json:"work"`
	Other     []string `json:"other"`
}

type Phone struct {
	Mobile  []string `json:"mobile"`
	Work    []string `json:"work"`
	Home    []string `json:"home"`
	Other   []string `json:"other"`
	Secure  []string `json:"secure"`
}

type Network struct {
	Bitbucket  string `json:"bb"`
	Blog       string `json:"blog"`
	Facebook   string `json:"fb"`
	Github     string `json:"gh"`
	Gitlab     string `json:"gl"`
	Linkedin   string `json:"linkedin"`
	Other      string `json:"other"`
	Site       string `json:"site"`
	Twitter    string `json:"twitter"`
}

type Encryption struct {
	Keys []Key
}

type Key struct {
	PublicKey    string `json:"pubkey"`
	Fingerprint  string `json:"fp"`
}

type Contact struct {
	Name Name
	Addresses []Address
	Email Email
	Phone Phone
	Network Network
	Encryption Encryption
	Notes []string `json:"notes"`
}
