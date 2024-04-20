package gm_item

type AttackUsage struct {
	MinDamage uint    `json:"minDamage"`
	MaxDamage uint    `json:"maxDamage"`
	CoolDown  float64 `json:"coolDown"`
}

type EatingUsage struct {
	HealAmount uint
}
