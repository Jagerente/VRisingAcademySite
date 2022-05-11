import axios from "axios";
import { onMounted, ref } from "vue"
export function useItems(limit) {
    const items = ref([])
    const isItemsLoading = ref(true)
    const fetching = async () => {
        try {
            const response = await axios.get(
                "https://jsonplaceholder.typicode.com/posts",
                {
                    params: {
                        _page: 1,
                        _limit: limit,
                    },
                }
            );
            items.value = response.data;
        } catch (e) {
            alert("Error" + e);
        } finally {
            isItemsLoading.value = false;
        }
    }

    onMounted(fetching)

    return {
        items, isItemsLoading,
    }
}