<template>
    <div>
        <div class="w-100">
            <VueSlider
                class="mx-auto slider-blood"
                v-model="this.bloodQualityValue"
                :marks="this.marks"
                :tooltip-placement="'bottom'"
                :tooltip-formatter="this.bloodQualityValue === 0 ? ' ' : '{value}%'"
                :tooltip="'always'"
                :min="0"
                dragOnClick
            ></VueSlider>
        </div>
    </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import { useWindowSize } from 'vue-window-size';
import VueSlider from 'vue-slider-component';

export default {
    data() {
        return {
            value: 0,
            blood_colors: [
                '#000',
                '#350909',
                '#681212',
                '#9a1a1a',
                '#aa1d1d'
            ],
        }
    },
    components: {
        VueSlider
    },
    methods: {
        ...mapActions({
            updateBloodQuality: "bloodTypes/updateBloodQuality",
        }),
    },
    computed: {
        bloodQualityValue: {
            get() {
                return this.bloodQuality;
            },
            set(value) {
                // In case of lags
                // this.debounce(() => this.updateBloodQuality(value))
                this.updateBloodQuality(value)
            }
        },
        ...mapState({
            bloodQuality: (state) => state.bloodTypes.bloodQuality,
        }),
        marks() {
            const ysm = '-50px';
            const ymd = '-50';

            const marks = {
                0: {
                    label: 'I',
                    style: {
                        width: '8px',
                        height: '8px',
                        display: 'block',
                        backgroundColor: this.blood_colors[0],
                        transform: 'translate(-2px, -2px)'
                    },
                    labelStyle: {
                        transform: this.windowWidth > 576 ? 'translate(-5px, -55px)' : `translate(-4px, ${ysm})`,
                        font_size: '1rem'
                    },
                    labelActiveStyle: {
                        color: 'yellow'
                    }
                },
                30: {
                    label: 'II',
                    style: {
                        width: '8px',
                        height: '8px',
                        display: 'block',
                        backgroundColor: this.blood_colors[1],
                        transform: 'translate(-2px, -2px)'
                    },
                    labelStyle: {
                        transform: this.windowWidth > 576 ? 'translate(-8px, -55px)' : `translate(-5px, ${ysm})`,
                    },
                    labelActiveStyle: {
                        color: 'yellow'
                    }

                },
                60: {
                    label: 'III',
                    style: {
                        width: '8px',
                        height: '8px',
                        display: 'block',
                        backgroundColor: this.blood_colors[2],
                        transform: 'translate(-2px, -2px)'
                    },
                    labelStyle: {
                        transform: this.windowWidth > 576 ? 'translate(-11px, -55px)' : `translate(-8px, ${ysm})`,
                    },
                    labelActiveStyle: {
                        color: 'yellow'
                    }
                },
                90: {
                    label: 'IV',
                    style: {
                        width: '8px',
                        height: '8px',
                        display: 'block',
                        backgroundColor: this.blood_colors[3],
                        transform: 'translate(-2px, -2px)'
                    },
                    labelStyle: {
                        transform: this.windowWidth > 576 ? 'translate(-10px, -55px)' : `translate(-8px, ${ysm})`,
                    },
                    labelActiveStyle: {
                        color: 'yellow'
                    }
                },
                100: {
                    label: 'V',
                    style: {
                        width: '8px',
                        height: '8px',
                        display: 'block',
                        backgroundColor: this.blood_colors[4],
                        transform: 'translate(-2px, -2px)'
                    },
                    labelStyle: {
                        transform: this.windowWidth > 576 ? 'translate(-9px, -55px)' : `translate(-6px, ${ysm})`,
                    },
                    labelActiveStyle: {
                        color: 'yellow'
                    }
                },
            };
            return marks;
        },

    },
    setup() {
        const { width, height } = useWindowSize();
        function createDebounce() {
            let timeout = null;
            return function (fnc, delayMs) {
                clearTimeout(timeout);
                timeout = setTimeout(() => {
                    fnc();
                }, delayMs || 100);
            };
        }

        return {
            debounce: createDebounce(),
            windowWidth: width,
            windowHeight: height,
        };
    }
}
</script>

<style lang="scss" scoped>
@import "@/assets/styles/va_variables.scss";
@import "@/assets/styles/va_slider.scss";
</style>