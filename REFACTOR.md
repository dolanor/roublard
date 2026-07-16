# Refactor

the refactor to do once I'm in sync with the original upstream project
# components

Health.CurrentHealth -> Current
Health.MaxHealth -> Max

- group weapon and armor in "item" type, with Name and description shared
- MeleeWeapon MinimumDamage -> Min, …


# combat_systems.go

```diff
-var a *ecs.QueryResult = nil
+var a *ecs.QueryResult
```

# level.go

doesn't do anything in our game engine
```diff
-levelHeight = gd.ScreenHeight - gd.UIHeight
```

# userlog_system.go

remove all the ebiten logic and replace with g3n logic

# level.go

// TODO: Change this to check for WALL, not blocked
IsOpaque()
