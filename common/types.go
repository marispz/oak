package common

type (
	GameConfig struct {
		WelcomeMessage string  `json:"welcome_message"`
		XpRate         float64 `json:"xp_rate"`
		Rarity         Rarity  `json:"rarity"`
	}

	Rarity struct {
		Common    RarityItems `json:"common"`
		Uncommon  RarityItems `json:"uncommon"`
		Rare      RarityItems `json:"rare"`
		Legendary RarityItems `json:"legendary"`
	}

	RarityItems struct {
		Chance float64 `json:"chance"`
		Items  []Item  `json:"items"`
	}

	Item struct {
		Name           string `json:"name"`
		Damage         int    `json:"damage,omitempty"`
		Defense        int    `json:"defense,omitempty"`
		Durability     int    `json:"durability"`
		SpecialAbility string `json:"special_ability,omitempty"`
	}
)
