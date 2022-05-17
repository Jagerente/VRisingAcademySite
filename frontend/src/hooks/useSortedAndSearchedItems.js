import { computed, ref } from "vue"

export function useSortedAndSearchedItems(sortedItems, searchQuery) {

    const sortedAndSearchedItems = computed(() => {
        return sortedItems.value.filter((item) =>
            item.title.toLowerCase().includes(searchQuery.value.toLowerCase())
        );
    });

    return {
        sortedAndSearchedItems
    }
}