import { computed, onMounted, ref } from "vue"

export function useSortedItems(items) {
    const sortedItems = computed(() => {
        return [...items.value].sort((item1, item2) => {
            return item1["title"]?.localeCompare(
                item2["title"]
            );
        });

    })

    return {
        sortedItems
    }
}