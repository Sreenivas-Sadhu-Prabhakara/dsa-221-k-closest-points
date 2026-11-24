# K Closest Points to Origin

> NeetCode 150 · Heap — squared distances (never sqrt), then take k from a deterministic sort; the heap and quickselect variants are trade-offs, not upgrades. Go, stdlib only, tested.

## Problem

Of n plane points, return the k nearest the origin by Euclidean distance.

## Approach

First insight: comparisons never need the square root — x²+y² orders identically to its root and stays in integers. Second: this is a top-k selection problem with three canonical tools, and *which wins depends on context*: a full sort (done here — simplest, cache-friendly, deterministic output), a size-k max-heap (O(n log k), the right call when n streams or k ≪ n — the exact shape of dsa-070), or quickselect (expected O(n), order-scrambled output — dsa-150). At sheet scale the sort is the honest default; the README of a library would say the same. Ties are broken by (x, y) so tests can compare exactly instead of set-comparing.

**Complexity:** O(n log n) time, O(n) copy.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
