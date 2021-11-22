package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")

	states := make(map[string]string);
	fmt.Println(states);
	states["WA"] = "Washington";
	states["OR"] = "Oregen";
	states["CA"] = "California";
	fmt.Println(states);

	california := states["CA"];
	fmt.Println(california);

	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println(states);

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v);
	}

	keys := make([]string, len(states));
	i := 0;
	for k := range states {
			keys[i] = k;
			i++;
	}
	fmt.Println(keys);
	sort.Strings(keys);
	fmt.Println(keys);

	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, states[k]);
	}

}
