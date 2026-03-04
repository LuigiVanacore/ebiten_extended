# Valutazione ebiten_extended (aggiornata)

**Data:** Febbraio 2025  
**Stato build:** ✅ `go build ./...` passa  
**Stato test:** ✅ `go test ./...` passa (package core + subpackages)

---

## 1. Correzioni completate

| Item | Stato |
|------|-------|
| updateTransform | ✅ Non muta più `parentGeoM`; position nella matrice locale |
| Typo minimun → minimum | ✅ Corretto in `math2D/range.go` e callers |
| README | ✅ Conflitti risolti, doc aggiornata, Sprite (non SpriteNode) |
| ResourceManager.AddImage | ✅ Restituisce `error` |
| Esempi | ✅ drawnShape, animation, event compilano e funzionano |
| Layers panic | ✅ `AddNodeToLayer`/`AddNodeToLayerF` restituiscono `error` |
| Camera | ✅ Origine top-left |
| Layer priority | ✅ Draw ordina per priorità |
| Test critici | ✅ updateTransform, layer priority, camera, Layers |
| Node2D.GetWorldPosition | ✅ Fix embedding: Node2D.AddChildren passa se stesso come parent |
| example/camera | ✅ `panic(err)` sostituito con `log.Fatal(err)` |
| Drawable.GetLayer | ✅ Usato in World.DrawNode per z-order tra sibling |
| Layers | ✅ Usato da World per draw order (layer index) |

---

## 2. Punteggi attuali

| Categoria      | Prima | Attuale | Note                         |
|----------------|-------|---------|------------------------------|
| Architettura   | 7/10  | **7.5/10**  | Buona modularità, duplicazioni da ridurre |
| Qualità codice | 6/10  | **7.5/10**  | Bug principali corretti       |
| Documentazione | 5/10  | **7/10**    | README pulito, doc.go ok      |
| Test           | 5/10  | **7.5/10**  | Test critici aggiunti         |
| Coerenza API   | 5/10  | **6.5/10**  | Esempi allineati, restano Layer vs Layers |
| Esempi/Build   | 4/10  | **7.5/10**  | Tutti compilanti              |

**Media attuale: ~7.2/10**

---

## 3. Punti forti

- **Scene graph**: Node/Node2D con transform locale e world, MarkDirty per caching
- **Collision**: Spatial hash, eventi Enter/Stay/Exit, maschere
- **Event**: `Event[T]` generico con Connect/Emit/Disconnect
- **Input**: Action system con RawInputButton, IsActionPressed/JustPressed
- **Modularità**: Package separati (math2D, transform, collision, input, event, tween, fsm)
- **Build/test**: Tutti i package compilano, test distribuiti su pacchetti core

---

## 4. Issue aperti e miglioramenti

### 4.1 Priorità alta

*(nessuna al momento)*

### 4.2 Priorità media

**Drawable.GetLayer** ✅ (usato per z-order sibling)  
- Obbligatorio nell’interfaccia ma non usato da `World.DrawNode`  
- Azione: Usarlo per ordinare i fratelli nello stesso layer, oppure rimuoverlo

**Input.GetCursorPos**  
- Doc indica "World mapped" ma restituisce coordinate schermo  
- Azione: Mappare con la camera o chiarire la doc

**FSM vs stateMachine** ✅ 
- Due pacchetti state machine  
- Azione: Scegliere un’API principale e riscrivere in Go Generics

### 4.3 Priorità bassa

- **Scale in updateTransform**: `Transform.GetScale()` non usato nel rendering
- **Tween**: Nessun aggancio al game loop
- **Tilemap**: Solo struct dati, niente rendering
- **Collision**: `IsColliding` muta le shape (UpdateTransform)

---

## 5. Test coverage (stima)

| Package      | Test files          | Copertura stimata |
|-------------|---------------------|-------------------|
| main        | engine, node, node2D, world, camera, layers | ~60% |
| transform   | transform_test      | ~80%              |
| event       | event_test          | ~70%              |
| collision   | collisionEvents     | ~50%              |
| input       | inputManager        | ~50%              |
| math2D      | vector2D, range (via altri) | ~40%        |
| utils       | stack, byteset      | ~70%              |

**Suggerimento:** Introdurre `go test -cover` in CI e fissare una soglia minima.

---

## 6. Roadmap residuale

### Fase 1 — Robustezza (completata)

1. ~~Indagare e sistemare `Node2D.GetWorldPosition`~~ ✅
2. ~~Sostituire `panic` con `log.Fatal` in `example/camera`~~ ✅

### Fase 2 — Coerenza (2 giorni)

3. Input: cursor in world coords o doc aggiornata  
4. ~~Drawable.GetLayer~~ ✅  
5. ~~Deprecare Layers~~ ✅  
6. ~~Unificare FSM e stateMachine~~ ✅

### Fase 3 — Estensioni (3–5 giorni)

7. Scale in `updateTransform`  
8. Tween collegato al game loop  
9. Tilemap rendering  
10. Collision: evitare mutazione in-place

---

## 7. Riepilogo

Il progetto è in buono stato per uno sviluppo attivo di giochi 2D con Ebiten. I punti critici (updateTransform, Layer, AddImage, esempi, test principali) sono stati risolti. Restano da affrontare soprattutto la coerenza API (Layer/Layers, FSM, Drawable) e l’investigazione su `GetWorldPosition`, con estensioni come scale, tween e tilemap in secondo piano.
