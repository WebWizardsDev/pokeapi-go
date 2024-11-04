package structs

// EvolutionDetail represents the details required for a Pokémon to evolve.
type EvolutionDetail struct {
	Gender                interface{} `json:"gender"`
	HeldItem              interface{} `json:"held_item"`
	Item                  interface{} `json:"item"`
	KnownMove             interface{} `json:"known_move"`
	KnownMoveType         interface{} `json:"known_move_type"`
	Location              interface{} `json:"location"`
	MinAffection          interface{} `json:"min_affection"`
	MinBeauty             interface{} `json:"min_beauty"`
	MinHappiness          interface{} `json:"min_happiness"`
	MinLevel              int         `json:"min_level"`
	NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
	PartySpecies          interface{} `json:"party_species"`
	PartyType             interface{} `json:"party_type"`
	RelativePhysicalStats interface{} `json:"relative_physical_stats"`
	TimeOfDay             string      `json:"time_of_day"`
	TradeSpecies          interface{} `json:"trade_species"`
	Trigger               struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"trigger"`
	TurnUpsideDown bool `json:"turn_upside_down"`
}

// ChainLink represents a link in a Pokémon evolution chain. Each link holds details about evolution criteria,
// subsequent evolutions, and species information.
type ChainLink struct {
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo        []ChainLink       `json:"evolves_to"`
	IsBaby           bool              `json:"is_baby"`
	Species          struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
}

// EvolutionChain represents a chain of Pokémon evolutions, including details about baby Pokémon and evolution triggers.
type EvolutionChain struct {
	BabyTriggerItem interface{} `json:"baby_trigger_item"`
	Chain           *ChainLink  `json:"chain,omitempty"`
	ID              int         `json:"id"`
}

// EvolutionTrigger is a single evolution trigger.
type EvolutionTrigger struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}
