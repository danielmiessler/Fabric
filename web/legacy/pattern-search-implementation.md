# Pattern Search Implementation Plan

## Component Changes (PatternList.svelte)

### 1. Add Search Input
```svelte
<div class="px-4 pb-4 flex gap-4 items-center">
  <!-- Existing sort options -->
  <div class="flex-1"> <!-- Add flex-1 to push search to right -->
    <label class="flex items-center gap-2 text-sm text-muted-foreground">
      <input type="radio" bind:group={sortBy} value="alphabetical">
      Alphabetical
    </label>
    <label class="flex items-center gap-2 text-sm text-muted-foreground">
      <input type="radio" bind:group={sortBy} value="favorites">
      Favorites First
    </label>
  </div>
  <!-- New search input -->
  <div class="w-48"> <!-- Fixed width for search -->
    <Input
      type="text"
      bind:value={searchText}
      placeholder="Search patterns..."
    />
  </div>
</div>
```

### 2. Add Search Logic
```typescript
// Add to script section
let searchText = ""; // For pattern filtering

// Modify sortedPatterns to include search
$: filteredPatterns = patterns.filter(p => 
  p.patternName.toLowerCase().includes(searchText.toLowerCase())
);

$: sortedPatterns = sortBy === 'alphabetical'
  ? [...filteredPatterns].sort((a, b) => a.patternName.localeCompare(b.patternName))
  : [
      ...filteredPatterns.filter(p => $favorites.includes(p.patternName)).sort((a, b) => a.patternName.localeCompare(b.patternName)),
      ...filteredPatterns.filter(p => !$favorites.includes(p.patternName)).sort((a, b) => a.patternName.localeCompare(b.patternName))
    ];
```

### 3. Reset Search on Selection
```typescript
// In pattern selection click handler
searchText = ""; // Reset search before closing modal
dispatch('select', pattern.patternName);
dispatch('close');
```

## Implementation Steps

1. Import Input component
```typescript
import { Input } from "$lib/components/ui/input";
```

2. Add searchText variable and filtering logic
3. Update template to include search input
4. Add reset logic in pattern selection handler
5. Test search functionality:
   - Partial matches work
   - Case-insensitive search
   - Search resets on selection
   - Layout maintains consistency

## Expected Behavior

- Search updates in real-time as user types
- Matches are case-insensitive
- Matches can be anywhere in pattern name
- Search box clears when pattern is selected
- Sort options (alphabetical/favorites) still work with filtered results
- Maintains existing modal layout and styling