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

<style lang="scss">
@import 'bootstrap/scss/_functions';
@import 'bootstrap/scss/_variables';
@import 'bootstrap/scss/_mixins';

@include media-breakpoint-down(sm) {
	.vue-slider-mark-label {
		font-size: 16px;
	}
}

@include media-breakpoint-up(sm) {
	.vue-slider-mark-label {
		font-size: 24px;
	}
}

/* component style */
.vue-slider-disabled .vue-slider-rail {
	background-color: #ccc;
}

.vue-slider-disabled .vue-slider-dot-handle {
	background-color: #666;
}

.vue-slider-disabled .vue-slider-process {
	background-color: #666;
}

.vue-slider-disabled .vue-slider-mark-step {
	background-color: #666;
}

.vue-slider-disabled .vue-slider-mark-step-active {
	background-color: #ccc;
}

/* rail style */
.vue-slider-rail {
	background-color: #ccc;
	border-radius: 15px;
}

/* process style */
.vue-slider-process {
	background: linear-gradient(0.25turn, #000, $primary);
	/* background-color: #6200ee; */
	border-radius: 15px;
}

/* mark style */
.vue-slider-mark {
	z-index: 4;
}

.vue-slider-mark-step {
	width: 100%;
	height: 100%;
	border-radius: 50%;
	background-color: $primary;
}

.vue-slider-mark-step-active {
	background-color: white;
}

.vue-slider-mark-label {
// font-size: 24px;
	white-space: nowrap;
}

/* dot style */
.vue-slider-dot-handle {
	/* cursor: pointer; */
	position: relative;
	width: 100%;
	height: 100%;
	border-radius: 50%;
	background-color: $primary;
	box-sizing: border-box;
}

.vue-slider-dot-handle::after {
	content: "";
	position: absolute;
	left: 50%;
	top: 50%;
	width: 200%;
	height: 200%;
	background-color: #aa1d1d35;
	border-radius: 50%;
	transform: translate(-50%, -50%) scale(0);
	z-index: -1;
	transition: transform 0.2s;
}

.vue-slider-dot-handle-focus::after {
	transform: translate(-50%, -50%) scale(1);
}

.vue-slider-dot-handle-disabled {
	cursor: not-allowed;
	background-color: #666 !important;
}

.vue-slider-dot-tooltip {
	visibility: visible;
}

.vue-slider-dot-tooltip-show .vue-slider-dot-tooltip-inner {
	opacity: 1;
}

.vue-slider-dot-tooltip-show .vue-slider-dot-tooltip-inner-top {
	transform: rotateZ(-45deg);
}

.vue-slider-dot-tooltip-show .vue-slider-dot-tooltip-inner-bottom {
	transform: rotateZ(135deg);
}

.vue-slider-dot-tooltip-show .vue-slider-dot-tooltip-inner-left {
	transform: rotateZ(-135deg);
}

.vue-slider-dot-tooltip-show .vue-slider-dot-tooltip-inner-right {
	transform: rotateZ(45deg);
}

.vue-slider-dot-tooltip-inner {
	border-radius: 50% 50% 50% 0px;
	background-color: #aa1d1d;
	opacity: 0;
	transition: transform 0.4s cubic-bezier(0.25, 0.8, 0.25, 1), opacity 0.2s linear;
}

.vue-slider-dot-tooltip-inner-top {
	transform: translate(0, 50%) scale(0.01) rotate(-45deg);
}

.vue-slider-dot-tooltip-inner-bottom {
	transform: translate(0, -50%) scale(0.01) rotateZ(135deg);
}

.vue-slider-dot-tooltip-inner-left {
	transform: translate(50%, 0) scale(0.01) rotateZ(-135deg);
}

.vue-slider-dot-tooltip-inner-right {
	transform: translate(-50%, 0) scale(0.01) rotateZ(45deg);
}

.vue-slider-dot-tooltip-text {
	font-size: 12px;
	white-space: nowrap;
	text-align: center;
	color: #fff;
	width: 30px;
	height: 30px;
	display: flex;
	align-items: center;
	justify-content: center;
	box-sizing: content-box;
}

.vue-slider-dot-tooltip-inner-top .vue-slider-dot-tooltip-text {
	transform: rotateZ(45deg);
}

.vue-slider-dot-tooltip-inner-bottom .vue-slider-dot-tooltip-text {
	transform: rotateZ(-135deg);
}

.vue-slider-dot-tooltip-inner-left .vue-slider-dot-tooltip-text {
	transform: rotateZ(135deg);
}

.vue-slider-dot-tooltip-inner-right .vue-slider-dot-tooltip-text {
	transform: rotateZ(-45deg);
}

/*# sourceMappingURL=material.css.map */
</style>