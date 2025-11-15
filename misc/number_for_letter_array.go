package misc

// NumberForLetterArray retorna um slice de strings,
// onde cada posição contém uma letra a partir de 'a'.
// Exemplo: input = 3 → ["a", "b", "c"]

func NumberForLetterArray(input int) []string {
	out := make([]string, 0, input)

	for i := 0; i < input; i++ {
		out = append(out, string(rune(i+'a')))
	}

	return out
}
