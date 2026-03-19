# Cosa manca e cosa si può migliorare al progetto

Analisi dello stato attuale del progetto `ebiten_extended`, funzionalità mancanti rispetto alla ROADMAP e suggerimenti di miglioramento.

---

## 1. Funzionalità mancanti (rispetto alla ROADMAP)

### Alta priorità

- ~~**Transizioni integrate con SceneManager**~~ – ✅ *Risolto*: `SetTransitionDuration(d)` abilita fade su `PushScene`/`ReplaceScene`/`PopScene`. `SetTransitionColor` per il colore dell’overlay.

### Media priorità

- **ResourceManager**
  - Manca atlas texture (`LoadAtlas`, `GetAtlasRegion`)
  - Manca preload con callback di progresso
  - Manca hot-reload per sviluppo
- **Physics**
  - Nessun supporto per joint/constraint
  - Nessun trigger (corpi che non collidono fisicamente ma emettono eventi)
  - Layer/mask da formalizzare meglio
- **Tilemap**
  - Pathfinder non integrato con `TileMapNode`
  - Nessun supporto isometrico
  - Nessun tile animato
- **Input**
  - Nessun input buffering per combo (es. fighting games)
  - Nessun supporto touch esplicito per mobile

### Priorità bassa

- **Camera**: Nessun supporto multiple camere (minimap, split-screen)
- **UI**: Nessun `FocusManager` per navigazione tastiera
- **Debug**: Nessun `DebugOverlay` per FPS, memoria, draw calls

---

## 2. Miglioramenti tecnici

### Test e coverage

- **Coverage attuale**: ~36% sul core, ~18% UI, ~21% tilemap → aumentare la coverage
- **Package senza test**: `particles` (0%), `example/*` (0%)
- Aggiungere test per: `Sprite`, `AnimationPlayer`, `ResourceManager`, integrazione `SceneManager`/`Transition`

### Layout UI (`ui/layout.go`)

- `HBoxLayout` e `VBoxLayout` base avanzano di `Spacing` fisso senza considerare la size dei figli
- Solo `HBoxLayoutEx` / `VBoxLayoutEx` usano `SizeProvider`, ma `PanelNode` e altri componenti non implementano `SizeProvider`
- Manca un sistema di `Anchor` (top-left, center, stretch) per posizionamento relativo

### Transition e SceneManager

- `Transition` è standalone; va integrato con `PushScene`/`PopScene` (es. fade out → cambia scena → fade in)
- Nessun esempio che mostri `Transition` usato con `SceneManager`

### Documentazione

- Il `doc.go` principale non menziona particles, pooling, scene manager, transizioni
- Mancano tutorial o guida "Getting started" passo-passo
- Sistema layer/mask della collision da documentare meglio

### CI

- Il workflow esegue `go tool cover` ma non pubblica il report (es. Codecov)
- Non c'è un job dedicato per gli esempi (`example/*`)

---

## 3. Qualità del codice

### Punti di forza

- Engine con `Pause`, `TimeScale`, `SceneManager` già implementati
- Object pooling: `utils.Pool[T]` e `SpritePool` funzionanti
- Particle system: `ParticleEmitter` + `ParticleEmitterNode` presenti
- `ActionMap` con `SetKeyBinding` e `GetKeyForAction` già implementati
- Layout UI: `HBoxLayout`, `VBoxLayout`, `LayoutPanelNode` presenti
- CI con build, test `-race`, golangci-lint

### Possibili miglioramenti

- **`ParticleEmitterNode`**: verificare che il `Draw` sia completo e funzionale
- **Layout**: far implementare `SizeProvider` ai componenti UI principali, oppure usare size fisse nei layout base
- **`doc.go`**: aggiornare per includere particles, pooling, scene manager, transizioni

---

## 4. Priorità suggerite

| Priorità | Azione                                           | Effort stimato |
|----------|---------------------------------------------------|----------------|
| Alta     | Integrare `Transition` con `SceneManager`         | 1–2 giorni     |
| Alta     | Aggiungere test per `particles`, `spritePool`, `SceneManager` | 1 giorno       |
| Media    | Verificare/completare draw di `ParticleEmitterNode` | ½ giorno    |
| Media    | Far implementare `SizeProvider` ai componenti UI  | 1 giorno       |
| Media    | Documentare collision layers/mask                 | ½ giorno       |
| Bassa    | Integrare pathfinder con `TileMapNode`            | 1 giorno       |
| Bassa    | Aggiungere report coverage in CI (es. Codecov)    | ½ giorno       |

---

*Report generato il 19 marzo 2025.*
