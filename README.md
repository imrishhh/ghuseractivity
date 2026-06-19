# Github User Activity

A CLI based project to track user activity on GitHub. Being specific
just displays last 300 events (within 30 days) of users on GitHub.

## Building & Running

**With Docker**:

```bash
docker compose up --build

```

Then find the executable in `./build` folder, move the executable to project folder then run:

```bash
./ghuseractivity <username> <page>(Default=1)
```

**Example**:

```bash
./ghuseractivity torvalds 5
```

**Output**:

```
PUSHED changes to `torvalds/linux` - 21d ago
PUSHED changes to `torvalds/linux` - 23d ago
PUSHED changes to `torvalds/linux` - 23d ago
PUSHED changes to `torvalds/linux` - 24d ago
CREATED a issue at `torvalds/GuitarPedal` - 25d ago
PUSHED changes to `torvalds/GuitarPedal` - 25d ago
PUSHED changes to `torvalds/linux` - 25d ago
PUSHED changes to `torvalds/linux` - 25d ago
PUSHED changes to `torvalds/GuitarPedal` - 25d ago
PUSHED changes to `torvalds/linux` - 26d ago
PUSHED changes to `torvalds/GuitarPedal` - 26d ago
PUSHED changes to `torvalds/linux` - 26d ago
PUSHED changes to `torvalds/GuitarPedal` - 26d ago
PUSHED changes to `torvalds/linux` - 27d ago
PUSHED changes to `torvalds/linux` - 27d ago
PUSHED changes to `torvalds/GuitarPedal` - 27d ago
PUSHED changes to `torvalds/linux` - 27d ago
PUSHED changes to `torvalds/linux` - 27d ago
PUSHED changes to `torvalds/GuitarPedal` - 27d ago
PUSHED changes to `torvalds/linux` - 27d ago
OPENED a issue at `torvalds/GuitarPedal` - 27d ago
OPENED a issue at `torvalds/GuitarPedal` - 27d ago
PUSHED changes to `torvalds/linux` - 28d ago
CLOSED a issue at `torvalds/GuitarPedal` - 28d ago
PUSHED changes to `torvalds/linux` - 28d ago
PUSHED changes to `torvalds/linux` - 28d ago
PUSHED changes to `torvalds/linux` - 29d ago
PUSHED changes to `torvalds/linux` - 29d ago
```
