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
./ghuseractivity torvalds 1
```

**Output**:

```
pushed changes to torvalds/linux - 22m ago
pushed changes to torvalds/linux - 2h ago
pushed changes to torvalds/linux - 3h ago
pushed changes to torvalds/linux - 23h ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 1d ago
pushed changes to torvalds/linux - 2d ago
pushed changes to torvalds/linux - 2d ago
pushed changes to torvalds/linux - 2d ago
pushed changes to torvalds/linux - 2d ago
pushed changes to torvalds/linux - 2d ago
pushed changes to torvalds/linux - 3d ago
pushed changes to torvalds/linux - 3d ago
pushed changes to torvalds/linux - 4d ago
pushed changes to torvalds/GuitarPedal - 4d ago
pushed changes to torvalds/linux - 4d ago
pushed changes to torvalds/GuitarPedal - 4d ago
pushed changes to torvalds/linux - 4d ago
pushed changes to torvalds/linux - 4d ago
pushed changes to torvalds/linux - 4d ago
pushed changes to torvalds/GuitarPedal - 5d ago
```
