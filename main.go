package main

import (
	"fmt"
)

func main() {
	// arrays
	// {
	// 	a := [3]int{41, 42, 43}
	// 	fmt.Printf("%#v\n", a)

	// 	b := [...]string{"hello", "my", "friends"}
	// 	fmt.Printf("%v", b)

	// 	var c [2]int
	// 	c[0] = 1
	// 	c[1] = 2
	// 	fmt.Println(c)

	// 	fmt.Printf("a is of type %T\tbut the length is %v\n", a, len(a))
	// 	fmt.Printf("b is of type %T\tbut the length is %v\n", b, len(b))
	// 	fmt.Printf("c is of type %T\tbut the length is %v\n", c, len(c))

	// 	myArray := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	// 	fmt.Printf("%T", myArray)
	// }

	// {
	// 	xi := [8]int{10, 20, 30, 40}
	// 	for i, v := range xi {
	// 		fmt.Printf("at index %v: %v\n", i, v)
	// 	}
	// 	fmt.Printf("%v\n", len(xi[2:]))
	// 	// xi = append(xi, 50, 60)
	// 	fmt.Printf("%v\n", cap(xi[2:]))

	// }

	// delete from slice
	// {
	// 	fmt.Println("before change-------------")
	// 	xi:=[]int{00,10,20,30,40,50,60,70}
	// 	fmt.Println(xi)
	// 	fmt.Printf("the content of xi2 is: %v\n",xi)
	// 	fmt.Printf("the lenght is: %v\n",len(xi))
	// 	fmt.Printf("the initial capacity is: %v\n",cap(xi))
	// 	// remove 40
	// 	fmt.Println("after change-------------")
	// 	// xi=append(xi[:4], xi[5:]...)
	// 	xi = append(xi, 80,90)
	// 	fmt.Println(xi)
	// 	fmt.Printf("the content of xi2 is: %v\n",xi)
	// 	fmt.Printf("the lenght is: %v\n",len(xi))
	// 	fmt.Printf("the initial capacity is: %v\n",cap(xi))
	// 	fmt.Println("xi2-------------")
	// 	xi2:=make([]int,2,10)
	// 	fmt.Printf("the content of xi2 is: %v\n",xi2)
	// 	fmt.Printf("the lenght is: %v\n",len(xi2))
	// 	fmt.Printf("the initial capacity is: %v\n",cap(xi2))
	// }
		// multiplane slice
	// {
	// 	xc:= []string{"black","white","red","green"}
	// 	xf:=[]string{"mango","fufu","pizza","balango"}

	// 	xxw:=[][]string{xc,xf}
	// 	fmt.Println(xxw)
	// }
	// unsafe pointer of lice
	// {
	// 	a:= []int{0,1,2,3,4,5}
	// 	// a and b will have the same underlying array
	// 	// b:=a
	// 	// even with a slice it still piitns to the original array
	// 	b:=a[:4]
	// 	fmt.Printf("a is %v\t the len is %v the cap is %v\n",a,len(a),cap(a))
	// 	fmt.Printf("b is %v\t\t the len is %v the cap is %v\n",b,len(b),cap(b))
		
	// 	a[0]=7
	// 	fmt.Printf("a is %v\n",a)
	// 	fmt.Printf("b is %v\n",b)
	// 	a = append(a, 10)
	// 	a[0]=8
	// 	fmt.Printf("a is %v\n",a)
	// 	fmt.Printf("b is %v\n",b)
	// 	// when u absolutely want the two to
	// 	// to point to the same array then use an array
	// 	// or just use a pointer to the slice

	// 	// this new slice c has its own underlying array
	// 	c:=make([]int,len(a))
	// 	copy(c,a)
	// }

	// 
	
	// maps
	// {
	// 	am:=make(map[string]int)
	// 	am["moh"]=21
	// 	am["avd"]=18
	// 	fmt.Println(am)
	// 	delete(am, "avd")
	// 	fmt.Println(am["avd"])
	// 	if v,ok:=am["avd"];ok{
	// 		fmt.Println(v)
	// 	}else {
	// 		fmt.Println("no value")
	// 	}
	// }


	// use maps to count frequency of words in a book
	// {
	// 	// msxs:=make(map[string][]string) //this is safer
	// 	msxs:=map[string][]string{
	// 		`bond_james`:{`shaken, not stirred`, `martinis`, `fast cars`},
	// 		`moneypenny_jenny`:{`intelligence`, `literature`, `computer science`},
	// 		`no_dr`:{`cats`, `ice cream`, `sunsets`	},
	// 	}
	// 	for k,v:= range msxs{
	// 		fmt.Printf("%v likes\n",k)
	// 		for i,w:= range v{
	// 			fmt.Printf("\t\t\t%v:%v\n",i+1,w)
	// 		}
	// 	}
	// 	delete(msxs,"bond_james")
	// 	for k,v:= range msxs{
	// 		fmt.Printf("%v likes\n",k)
	// 		for i,w:= range v{
	// 			fmt.Printf("\t\t\t%v:%v\n",i+1,w)
	// 		}
	// 	}
	// }

	{
		words:=make([]string,0,1000)
		words = append(words, "in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some", "advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever", "you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all", "the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve", "had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually", "communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant", "a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to", "reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious", "natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few", "veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach", "itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so", "it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a", "politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown", "men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have", "feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by", "some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on", "the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least", "the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and", "marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of", "infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i", "forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly", "repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out", "unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i", "come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be", "founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain", "point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from", "the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in", "uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no", "more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.", "only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was", "exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which", "i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of", "successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,", "some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were", "related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes", "ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with", "that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the", "creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic", "readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and", "which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out", "all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust", "floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my", "interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my", "family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western", "city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,", "and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of", "buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s","brother,", "who", "came", "here", "in", "fifty-one,", "sent", "a", "substitute", "to", "the", "civil", "war,", "and", "started", "the", "wholesale", "hardware", "business", "that", "my", "father", "carries", "on", "today.", "i", "never", "saw", "this", "great-uncle,", "but", "i’m", "supposed", "to", "look", "like", "him—with", "special", "reference", "to", "the", "rather", "hard-boiled", "painting", "that", "hangs", "in", "father’s", "office.", "i", "graduated", "from", "new", "haven", "in", "1915,", "just", "a", "quarter", "of", "a", "century", "after", "my", "father,", "and", "a", "little", "later", "i", "participated", "in", "that", "delayed", "teutonic", "migration", "known", "as", "the", "great", "war.", "i", "enjoyed", "the", "counter-raid", "so", "thoroughly", "that", "i", "came", "back", "restless.", "instead", "of", "being", "the", "warm", "centre", "of", "the", "world,", "the", "middle", "west", "now", "seemed", "like", "the", "ragged", "edge", "of", "the", "universe—so", "i", "decided", "to", "go", "east", "and", "learn", "the", "bond", "business.", "everybody", "i", "knew", "was", "in", "the", "bond", "business,", "so", "i", "supposed", "it", "could", "support", "one", "more", "single", "man.", "all", "my", "aunts", "and", "uncles", "talked", "it", "over", "as", "if", "they", "were", "choosing", "a", "prep", "school", "for", "me,", "and", "finally", "said,", "why—ye-es,", "with", "very", "grave,", "hesitant", "faces.", "father", "agreed", "to", "finance", "me", "for", "a", "year,", "and", "after", "various", "delays", "i", "came", "east,", "permanently,", "i", "thought,", "in", "the", "spring", "of", "twenty-two.", "the", "practical", "thing", "was", "to", "find", "rooms", "in", "the", "city,", "but", "it", "was", "a", "warm", "season,", "and", "i", "had", "just", "left", "a", "country", "of", "wide", "lawns", "and", "friendly", "trees,", "so", "when", "a", "young", "man", "at", "the", "office", "suggested", "that", "we", "take", "a", "house", "together", "in", "a", "commuting", "town,", "it", "sounded", "like", "a", "great", "idea.", "he", "found", "the", "house,", "a", "weather-beaten", "cardboard", "bungalow", "at", "eighty", "a", "month,", "but", "at", "the", "last", "minute", "the", "firm", "ordered", "him", "to", "washington,", "and", "i", "went", "out", "to", "the", "country", "alone.", "i", "had", "a", "dog—at", "least", "i", "had", "him", "for", "a", "few", "days", "until", "he", "ran", "away—and", "an", "old", "dodge", "and", "a", "finnish", "woman,", "who", "made", "my", "bed", "and", "cooked", "breakfast", "and", "muttered", "finnish", "wisdom", "to", "herself", "over", "the", "electric", "stove.", "it", "was", "lonely", "for", "a", "day", "or", "so", "until", "one", "morning", "some", "man,", "more", "recently", "arrived", "than", "i,", "stopped", "me", "on", "the", "road.", "how", "do", "you", "get", "to", "west", "egg", "village?", "he", "asked", "helplessly.", "i", "told", "him.", "and", "as", "i", "walked", "on", "i", "was", "lonely", "no", "longer.", "i", "was", "a", "guide,", "a", "pathfinder,", "an", "original", "settler.", "he", "had", "casually", "conferred", "on", "me", "the", "freedom", "of", "the", "neighbourhood.", "and", "so", "with", "the", "sunshine", "and", "the", "great", "bursts", "of", "leaves", "growing", "on", "the", "trees,", "just", "as", "things", "grow", "in", "fast", "movies,", "i", "had", "that", "familiar", "conviction", "that", "life", "was", "beginning", "over", "again", "with", "the", "summer.", "there", "was", "so", "much", "to", "read,", "for", "one", "thing,", "and", "so", "much", "fine", "health", "to", "be", "pulled", "down", "out", "of", "the", "young", "breath-giving", "air.", "i", "bought", "a", "dozen", "volumes", "on", "banking", "and", "credit", "and", "investment", "securities,", "and", "they", "stood", "on", "my", "shelf", "in", "red")
		count:=make(map[string]int)
		for _,w:=range words{
			count[w]++
		}
		// fmt.Println(count)
		for k,v:=range count{
			fmt.Println(k,"\t\t\t\t",v)
		}
	}
}
