# Collision Layers and Masks

Questo documento descrive il sistema di layer e mask per il filtraggio delle collisioni nel package `collision`.

## Concetto

Ogni `Collider` e `Area2D` possiede una `CollisionMask` che determina con quali altri oggetti può collidere. La mask è composta da due insiemi di bit (`utils.ByteSet`):

- **Identity**: il layer (o i layer) a cui appartiene l'oggetto — *cosa è*
- **Mask**: i layer con cui l'oggetto può collidere — *con chi risponde*

## Regola di collisione

Una coppia (A, B) collide **solo se entrambi** la consentono:

1. `A.mask` deve includere `B.identity` (A vuole collidere con il layer di B)
2. `B.mask` deve includere `A.identity` (B vuole collidere con il layer di A)

Il `CollisionManager` verifica entrambe le condizioni prima di emettere eventi Enter/Stay/Exit.

## Layer predefiniti

I preset definiscono layer come potenze di 2:

| Costante      | Valore | Uso tipico        |
|---------------|--------|-------------------|
| LayerPlayer   | 1<<0   | Personaggio       |
| LayerEnemy    | 1<<1   | Nemici            |
| LayerWorld    | 1<<2   | Terreno, muri     |
| LayerPickup   | 1<<3   | Oggetti raccoglibili |
| LayerProjectile | 1<<4 | Proiettili        |

## Preset pronti all'uso

| Preset       | Identity   | Collide con              |
|--------------|------------|---------------------------|
| MaskPlayer   | LayerPlayer | World, Pickup            |
| MaskEnemy    | LayerEnemy  | World, Player            |
| MaskWorld    | LayerWorld  | Player, Enemy, Projectile |
| MaskPickup   | LayerPickup | Player                    |
| MaskProjectile | LayerProjectile | World, Enemy       |

## Matrice delle collisioni (preset)

```
              Player  Enemy   World   Pickup  Projectile
Player          -      ✓       ✓       ✓         -
Enemy           ✓      -       ✓       -         ✓
World           ✓      ✓       -       -         ✓
Pickup          ✓      -       -       -         -
Projectile      -      ✓       ✓       -         -
```

## Utilizzo

### Con i preset

```go
playerCol, _ := collision.NewCollider("player", shape, collision.MaskPlayer)
worldCol, _ := collision.NewCollider("wall", wallShape, collision.MaskWorld)
```

### Con layer custom

```go
const (
    LayerTrap    utils.ByteSet = 1 << 5
    LayerSensor  utils.ByteSet = 1 << 6
)

// Trappola: collide solo con Player
trapMask := collision.NewPresetMask(LayerTrap, collision.LayerPlayer)

// Sensore: collide con Player e Enemy (per trigger area)
sensorMask := collision.NewPresetMask(LayerSensor, collision.LayerPlayer, collision.LayerEnemy)
```

### Maschere manuali

```go
identity := utils.ByteSet(1)        // layer 1
mask := utils.ByteSet(1 | 2 | 4)    // collide con layer 1, 2, 4
cm := collision.NewCollisionMask(identity, mask)
```

## Note

- Usare sempre potenze di 2 (1<<0, 1<<1, …) per evitare sovrapposizioni tra layer
- Un oggetto può appartenere a più layer: `identity = LayerPlayer | LayerSensor`
- Modificare la mask a runtime con `SetCollisionMask` (es. invincibilità: mask vuota temporaneamente)
