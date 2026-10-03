package handler

import (
	"net/url"
	"testing"
)

// TestIntQuery: intQuery ফাংশন ঠিকমতো URL থেকে সংখ্যা পড়ে কিনা যাচাই করে।
// এটা একটা "table-driven test": অনেকগুলো case একটা টেবিলে রেখে
// একই কোড দিয়ে সবগুলো একে একে চালাই।
func TestIntQuery(t *testing.T) {
	tests := []struct {
		name    string // এই case-এর নাম (ফেল করলে এটা দেখাবে)
		rawURL  string // URL-এর ? চিহ্নের পরের অংশ
		def     int    // কিছু না দিলে কোন মান ফেরত আসবে
		want    int    // আমরা কোন মান আশা করছি
		wantErr bool   // error আশা করছি কিনা
	}{
		{name: "সংখ্যা দিলে সেটাই ফেরত দেয়", rawURL: "page=3", def: 1, want: 3},
		{name: "কিছু না দিলে default দেয়", rawURL: "", def: 1, want: 1},
		{name: "খালি মান দিলেও default দেয়", rawURL: "page=", def: 7, want: 7},
		{name: "ঋণাত্মক সংখ্যাও পড়ে", rawURL: "page=-5", def: 1, want: -5},
		{name: "অক্ষর দিলে error", rawURL: "page=abc", def: 1, wantErr: true},
		{name: "দশমিক দিলে error", rawURL: "page=2.5", def: 1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ১. প্রস্তুতি (Arrange): টেক্সট থেকে url.Values বানাও
			q, err := url.ParseQuery(tt.rawURL)
			if err != nil {
				t.Fatalf("test-এর setup-এই ভুল: %v", err)
			}

			// ২. কাজ (Act): যে ফাংশন পরীক্ষা করছি সেটা চালাও
			got, err := intQuery(q, "page", tt.def)

			// ৩. যাচাই (Assert): ফলাফল মিলিয়ে দেখো
			if tt.wantErr {
				if err == nil {
					t.Fatalf("error আশা করেছিলাম, কিন্তু পেলাম %d", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("error আশা করিনি, কিন্তু পেলাম: %v", err)
			}
			if got != tt.want {
				t.Errorf("intQuery() = %d, আশা করেছিলাম %d", got, tt.want)
			}
		})
	}
}

// TestBoolQuery: boolQuery ফাংশন URL থেকে true/false ঠিকমতো পড়ে কিনা যাচাই করে।
// boolQuery ফেরত দেয় *bool, তাই তিন রকম ফল হতে পারে:
// nil (কিছু দেওয়া হয়নি), true, অথবা false।
func TestBoolQuery(t *testing.T) {
	tests := []struct {
		name    string // এই case-এর নাম
		rawURL  string // URL-এর ? চিহ্নের পরের অংশ
		wantNil bool   // nil আশা করছি কিনা
		want    bool   // nil না হলে ভেতরে কোন মান আশা করছি
		wantErr bool   // error আশা করছি কিনা
	}{
		{name: "true দিলে true", rawURL: "done=true", want: true},
		{name: "false দিলে false", rawURL: "done=false", want: false},
		{name: "কিছু না দিলে nil", rawURL: "", wantNil: true},
		{name: "খালি মান দিলেও nil", rawURL: "done=", wantNil: true},
		{name: "1 দিলেও true (অবাক করা!)", rawURL: "done=1", want: true},
		{name: "0 দিলে false", rawURL: "done=0", want: false},
		{name: "t দিলেও true", rawURL: "done=t", want: true},
		{name: "TRUE বড় হাতে দিলেও true", rawURL: "done=TRUE", want: true},
		{name: "yes দিলে error", rawURL: "done=yes", wantErr: true},
		{name: "উল্টাপাল্টা লিখলে error", rawURL: "done=abc", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ১. প্রস্তুতি (Arrange)
			q, err := url.ParseQuery(tt.rawURL)
			if err != nil {
				t.Fatalf("test-এর setup-এই ভুল: %v", err)
			}

			// ২. কাজ (Act): শুধু boolQuery চালাও, নিজে কিছু parse করো না
			got, err := boolQuery(q, "done")

			// ৩. যাচাই (Assert)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("error আশা করেছিলাম, কিন্তু পাইনি")
				}
				return
			}
			if err != nil {
				t.Fatalf("error আশা করিনি, কিন্তু পেলাম: %v", err)
			}

			// বাক্স থাকার কথাই না?
			if tt.wantNil {
				if got != nil {
					t.Fatalf("nil আশা করেছিলাম, কিন্তু পেলাম %v", *got)
				}
				return
			}

			// বাক্স থাকার কথা, আগে দেখো বাক্সটা আছে কিনা
			if got == nil {
				t.Fatalf("একটা মান আশা করেছিলাম, কিন্তু পেলাম nil")
			}

			// বাক্স খুলে ভেতরের মান মিলিয়ে দেখো
			if *got != tt.want {
				t.Errorf("boolQuery() = %v, আশা করেছিলাম %v", *got, tt.want)
			}
		})
	}
}
