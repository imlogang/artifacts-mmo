package resource

type Response struct {
	Data Data `json:"data"`
}
type Data struct {
	Character       `json:",inline,omitempty"`
	NestedCharacter Character   `json:"character,omitempty"`
	Cooldown        Cooldown    `json:"cooldown,omitempty"`
	Destination     Destination `json:"destination,omitempty"`
	Fight           Fight       `json:"fight,omitempty"`
	HPRestored      int         `json:"hp_restored,omitempty"`
}
type Character struct {
	Name                    string `json:"name"`
	Skin                    string `json:"skin"`
	Level                   int    `json:"level"`
	Xp                      int    `json:"xp"`
	MapXp                   int    `json:"max_xp"`
	TotalXp                 int    `json:"total_xp"`
	Gold                    int    `json:"gold"`
	Speed                   int    `json:"speed"`
	MiningLevel             int    `json:"mining_level"`
	MiningXP                int    `json:"mining_xp"`
	MiningMaxXP             int    `json:"mining_max_xp"`
	WoodCuttingLevel        int    `json:"woodcutting_level"`
	WoodCuttingXP           int    `json:"woodcutting_xp"`
	WoodCuttingMaxXP        int    `json:"woodcutting_max_xp"`
	FishingLevel            int    `json:"fishing_level"`
	FishingXP               int    `json:"fishing_xp"`
	FishingMaxXP            int    `json:"fishing_max_xp"`
	WeaponCraftingLevel     int    `json:"weaponcrafting_level"`
	WeaponCraftingXP        int    `json:"weaponcrafting_xp"`
	WeaponCraftingMaxXP     int    `json:"weaponcrafting_max_xp"`
	GearCraftingLevel       int    `json:"gearcrafting_level"`
	GearCraftingXP          int    `json:"gearcrafting_xp"`
	GearCraftingMaxXP       int    `json:"gearcrafting_max_xp"`
	JewelryCraftingLevel    int    `json:"jewelrycrafting_level"`
	JewelryCraftingXP       int    `json:"jewelrycrafting_xp"`
	JewelryCraftingMaxXP    int    `json:"jewelrycrafting_max_xp"`
	CookingCraftingLevel    int    `json:"cookingcrafting_level"`
	CookingCraftingXP       int    `json:"cookingcrafting_xp"`
	CookingCraftingMaxXP    int    `json:"cookingcrafting_max_xp"`
	Hp                      int    `json:"hp"`
	Haste                   int    `json:"haste"`
	CriticalStrike          int    `json:"critical_strike"`
	Stamina                 int    `json:"stamina"`
	AttackEarth             int    `json:"attack_earth"`
	AttackWater             int    `json:"attack_water"`
	AttackFire              int    `json:"attack_fire"`
	AttackAir               int    `json:"attack_air"`
	DamageEarth             int    `json:"damage_earth"`
	DamageWater             int    `json:"damage_water"`
	DamageFire              int    `json:"damage_fire"`
	DamageAir               int    `json:"damage_air"`
	ResEarth                int    `json:"res_earth"`
	ResWater                int    `json:"res_water"`
	ResFire                 int    `json:"res_fire"`
	ResAir                  int    `json:"res_air"`
	XLoc                    int    `json:"x"`
	YLoc                    int    `json:"y"`
	Cooldown                int    `json:"cooldown,omitempty"`
	CooldownExpiration      string `json:"cooldown_expiration"`
	WeaponSlot              string `json:"weapon_slot"`
	ShieldSlot              string `json:"shield_slot"`
	HelmetSlot              string `json:"helmet_slot"`
	BodyArmorSlot           string `json:"body_armor_slot"`
	LegArmorSlot            string `json:"leg_armor_slot"`
	BootsSlot               string `json:"boots_armor_slot"`
	Ring1Slot               string `json:"ring1_slot"`
	Ring2Slot               string `json:"ring2_slot"`
	AmuletSlot              string `json:"amulet_slot"`
	Artifact1Slot           string `json:"artifact1_slot"`
	Artifact2Slot           string `json:"artifact2_slot"`
	Artifact3Slot           string `json:"artifact3_slot"`
	Consumable1Slot         string `json:"consumable1_slot"`
	Consumable1SlotQuantity int    `json:"consumable1_slot_quantity"`
	Consumable2Slot         string `json:"consumable2_slot"`
	Consumable2SlotQuantity int    `json:"consumable2_slot_quantity"`
	Task                    string `json:"task"`
	TaskType                string `json:"task_type"`
	TaskProgress            int    `json:"task_progress"`
	TaskTotal               int    `json:"task_total"`
	Inventory               []struct {
		Slot     int    `json:"slot"`
		Code     string `json:"code"`
		Quantity int    `json:"quantity"`
	}
	InventoryMax int `json:"inventory_max_items"`
}

type Cooldown struct {
	TotalSeconds     int    `json:"total_seconds"`
	RemainingSeconds int    `json:"remaining_seconds"`
	StartedAt        string `json:"started_at"`
	Expiration       string `json:"expiration"`
	Reason           string `json:"reason"`
}

type Destination struct {
	Name string `json:"name"`
	Skin string `json:"skin"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type Content struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

// Is this the right way to type this or should they all go first?
type Fight struct {
	XP                 int                `json:"xp"`
	Gold               int                `json:"gold"`
	Drops              Drops              `json:"drops"`
	Turns              int                `json:"turns"`
	MonsterBlockedHits MonsterBlockedHits `json:"monster_blocked_hits"`
	PlayerBlockedHits  PlayerBlockedHits  `json:"player_blocked_hits"`
	Logs               map[string]string  `json:"logs"`
	Result             string             `json:"result"`
}

type Drops struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type MonsterBlockedHits struct {
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Water int `json:"water"`
	Air   int `json:"air"`
	Total int `json:"total"`
}

type PlayerBlockedHits struct {
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Water int `json:"water"`
	Air   int `json:"air"`
	Total int `json:"total"`
}

type HPRestored struct {
	HPRestored int `json:"hp_restored"`
}
