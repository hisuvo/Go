Go-তে **Channels** হলো এমন একটি মাধ্যম যার সাহায্যে বিভিন্ন Goroutines নিজেদের মধ্যে ডেটা আদান-প্রদান (Communication) এবং সিনক্রোনাইজেশন (Synchronization) করতে পারে (Yuan et al., 2021)। গো-তে মূলত দুই ধরণের চ্যানেল আছে: **Unbuffered Channel** এবং **Buffered Channel** (Yuan et al., 2021)।

সহজ ভাষায় এবং বাস্তব উদাহরণের মাধ্যমে নিচে এদের বিস্তারিত আলোচনা করা হলো।

---

## ১. Unbuffered Channel (বাফারহীন চ্যানেল)

Unbuffered Channel-এর কোনো নিজস্ব মেমোরি ব্যাকআপ বা "সংরক্ষণাগার" থাকে না। এর ধারণক্ষমতা (Capacity) হলো **০ (শূণ্য)** (Saioc, 2025)।

- **কাজের ধরন:** এতে ডেটা পাঠাতে (Send) হলে একজন গ্রহণকারীকে (Receiver) অবশ্যই একই সময়ে উপস্থিত থাকতে হবে (Yuan et al., 2021)। যতক্ষণ না কোনো Goroutine ডেটা গ্রহণ করছে, ততক্ষণ ডেটা পাঠানো Goroutine-টি ব্লক (আটকে) হয়ে বসে থাকবে (Yuan et al., 2021)। একে **Synchronous** কমিউনিকেশন বলা হয়।
- **বাস্তব উদাহরণ:** রিলে রেস (Relay Race) বা লাঠি বদল খেলার মতো। একজন রানার লাঠি তখনই হাত থেকে ছাড়তে পারেন যখন অন্যজন সেটি হাত বাড়িয়ে ধরে।

### কোড উদাহরণ:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// একটি unbuffered channel তৈরি করা হলো
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine: ডেটা পাঠানোর চেষ্টা করছি...")
		ch <- "হ্যালো বাংলাদেশ!" // এখানে এসে Goroutine-টি ব্লক হবে, যতক্ষণ না main এটি গ্রহণ করছে
		fmt.Println("Goroutine: ডেটা পাঠানো সফল!")
	}()

	time.Sleep(2 * time.Second) // বুঝার সুবিধার জন্য ২ সেকেন্ড বিরতি
	fmt.Println("Main: ডেটা গ্রহণ করার জন্য প্রস্তুত।")

	msg := <-ch // ডেটা রিসিভ করা হলো
	fmt.Println("Main প্রাপ্ত ডেটা:", msg)
}

```

---

## ২. Buffered Channel (বাফারযুক্ত চ্যানেল)

Buffered Channel-এর নিজস্ব একটি নির্দিষ্ট মেমোরি পকেট বা কিউ (Queue) থাকে, যেখানে সে ডেটা সাময়িকভাবে জমা রাখতে পারে (Saioc, 2025)। এর ধারণক্ষমতা (Capacity) **০-এর বেশি** হয় (Saioc, 2025)।

- **কাজের ধরন:** বাফারের সাইজ যতক্ষণ না সম্পূর্ণ পূর্ণ (Full) হচ্ছে, ততক্ষণ ডেটা পাঠানো Goroutine-টি ব্লক হবে না; সে ডেটা রেখে নিজের কাজ চালিয়ে যেতে পারবে (Yuan et al., 2021)。 একে **Asynchronous** কমিউনিকেশন বলে (Yuan et al., 2021)। তবে বাফার ফুল হয়ে গেলে এটিও ব্লকিং আচরণ করবে (Yuan et al., 2021)。
- **বাস্তব উদাহরণ:** চিঠি এবং পোস্টবক্সের মতো। পোস্টম্যান বা আপনি এসে চিঠির বক্সে চিঠিটি ফেলে দিয়ে চলে যেতে পারেন (ব্লক হতে হয় না)। প্রাপক পরে তার সুবিধামতো এসে বক্স থেকে চিঠি নিয়ে যাবে। কিন্তু বক্সটি যদি চিঠিতে একদম ঠাসা (Full) থাকে, তবে খালি না হওয়া পর্যন্ত নতুন চিঠি ঢোকানো যাবে না।

### কোড উদাহরণ:

```go
package main

import "fmt"

func main() {
	// ২ ধারণক্ষমতার (Capacity) একটি buffered channel তৈরি করা হলো
	ch := make(chan string, 2)

	// কোনো রিসিভার ছাড়াই পরপর দুটি ডেটা পাঠানো যাবে, কোড ব্লক হবে না
	ch <- "প্রথম বার্তা"
	ch <- "দ্বিতীয় বার্তা"

	fmt.Println("বাফার ফুল না হওয়া পর্যন্ত কোড ব্লক হয়নি।")

	// এবার ডেটা রিসিভ করা হচ্ছে
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

```

---

## চিত্রগত তুলনা (Visual Comparison)

নিচের ধারণাগত আর্কিটেকচারটি লক্ষ্য করলে বিষয়টি আরও স্পষ্ট হবে:

```text
[ Unbuffered Channel ] -> Capacity: 0
Sender Goroutine  ========[ No Buffer ]========>  Receiver Goroutine
(উভয়কেই একই সময়ে লাইনে থাকতে হবে, নয়তো ব্লকিং ঘটবে)

[ Buffered Channel ]   -> Capacity: 2
Sender Goroutine  ===> [ Slot 1 ] [ Slot 2 ] ===> Receiver Goroutine
(বাফার খালি থাকলে Sender ডেটা রেখেই চলে যেতে পারে, রিসিভার পরে এসে নেয়)

```

---

## ৩. গুরুত্বপূর্ণ ইন্টারভিউ প্রশ্ন ও উত্তর (Interview Questions)

**প্রশ্ন ১: Unbuffered এবং Buffered চ্যানেলের মধ্যে মূল পার্থক্য কী?**

- **উত্তর:** মূল পার্থক্য হলো ধারণক্ষমতা (Capacity) এবং ব্লকিং মেকানিজম (Yuan et al., 2021; Saioc, 2025)। Unbuffered চ্যানেলের ক্যাপাসিটি ০, তাই এটি হ্যান্ডশেক (Handshake) বা সিনক্রোনাস পদ্ধতিতে কাজ করে (Yuan et al., 2021)। Buffered চ্যানেলের নির্দিষ্ট ক্যাপাসিটি থাকে এবং বাফার পূর্ণ না হওয়া পর্যন্ত এটি Sender-কে ব্লক করে না (Yuan et al., 2021)।

**প্রশ্ন ২: Deadlock (ডেডলক) কী এবং এটি Unbuffered চ্যানেলে কেন ঘটে?**

- **উত্তর:** যখন কোনো Goroutine একটি চ্যানেলে ডেটা পাঠায় কিন্তু সেই ডেটা গ্রহণ করার জন্য অন্য কোনো Goroutine সচল থাকে না, তখন পাঠানো Goroutine-টি চিরকালের জন্য ব্লক হয়ে যায় (Yuan et al., 2021; Saioc, 2025)। যদি অ্যাপ্লিকেশনের সব Goroutine এভাবে ব্লক হয়ে বসে থাকে, তখন Go Runtime একটি fatal error দেয়, একেই ডেডলক বলে (Saioc, 2025)।
- _উদাহরণ:_ `main` ফাংশনে কোনো আলাদা Goroutine ছাড়া সরাসরি `ch <- "data"` লিখলে কোড সেখানেই আটকে যাবে, কারণ নেওয়ার মতো কেউ নেই।

**প্রশ্ন ৩: একটি Closed (বন্ধ) চ্যানেল থেকে ডেটা রিড (Read) করলে কী ঘটবে?**

- **উত্তর:** চ্যানেল বন্ধ হয়ে গেলেও যদি বাফারে কোনো পুরাতন ডেটা জমা থাকে, তবে প্রথমে সেই ডেটাগুলো সফলভাবে রিড হবে (Saioc, 2025)। বাফার সম্পূর্ণ খালি হয়ে যাওয়ার পর (অথবা Unbuffered চ্যানেলের ক্ষেত্রে সরাসরি) রিড করলে চ্যানেলের ডেটা টাইপের **Zero Value** (যেমন: `int` হলে `0`, `string` হলে `""`) রিটার্ন করবে এবং কোনো ব্লকিং হবে না (Saioc, 2025)।

**প্রশ্ন ৪: একটি Closed চ্যানেলে ডেটা রাইট (Write) করতে গেলে কী হবে?**

- **উত্তর:** কোনো বন্ধ হয়ে যাওয়া চ্যানেলে পুনরায় ডেটা পাঠাতে (`ch <- data`) গেলে Go Runtime সরাসরি **Panic** করবে (Saioc, 2025)। তাই চ্যানেল বন্ধ করার দায়িত্ব সাধারণত Sender Goroutine-এরই হওয়া উচিত।

---

## References

- Yuan, T., Li, G., Lu, J., Liu, C., Li, L., & Xue, J. (2021). GoBench: A Benchmark Suite of Real-World Go Concurrency Bugs. _2021 IEEE/ACM International Symposium on Code Generation and Optimization (CGO)_, 187-199. [https://doi.org/10.1109/cgo51591.2021.9370317]()
- Saioc, G. V. (2025). Dynamic Partial Deadlock Detection and Recovery via Garbage Collection. _Department of Computer Science, Aarhus University_.

Cited by: 52 (Yuan et al., 2021)
