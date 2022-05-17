<template>
  <div class="container flex-column ">
    <div class="d-flex justify-content-center">
      <img class="image" style="margin-top: 0px;margin-bottom: -135px;" draggable="false" :src="
        require('@/assets/images/ui/main_page/VAlogo.webp')
      ">
    </div>
    <h1 class="text-center m-5">Welcome to the academy, V Academy!
      <br>This website is Work In Progress, and we are
      working
      hard to fill up the
      database with missing items and develop new features.
      <br>If you want to contribute or leave feedback you can do
      it in <a
        href="https://docs.google.com/forms/d/e/1FAIpQLSdNlAa_hyXcHsMTJy59alBP5hxhiWsebMHNr9M52He0s5mUsQ/viewform?usp=sf_link"
        target="blank">this form.</a>
    </h1>
    <div class="d-flex justify-content-center">
      <div class="mx-3">
        <div class="position-absolute orb1">
        </div>
        <a href="https://vk.com/vrisingacademy" target="blank">
          <img class=" image position-relative " :src="
            require('@/assets/images/ui/main_page/VK.webp')
          ">
        </a>
      </div>
      <div class="position-relative mx-3">
        <div class="position-absolute orb2">
        </div>
        <a href="https://t.me/vrisingacademy" target="blank">
          <img class="image position-relative" :src="
            require('@/assets/images/ui/main_page/TG.webp')
          ">
        </a>
      </div>
      <div class="position-relative mx-3">
        <div class="position-absolute orb3">
        </div>
        <a href="https://boosty.to/vrisingacademy" target="blank">
          <img class="image position-relative " type="image" :src="
            require('@/assets/images/ui/main_page/Boosty.webp')
          ">
        </a>
      </div>
    </div>
    <div class="d-flex justify-content-center" style="translate: 0px 150px">
      <p class="text-center">© 2022 V Rising Academy · Genshin Academy Team</p>
    </div>
  </div>
</template>

<script>
export default {
};
</script>

<style lang="scss" scoped>
* {
  user-select: none;
}

$violet1: violet;
$blue1: #0ff;
$violet2: #29a0dc;
$blue2: white;
$violet3: orange;
$blue3: white;
$size: 85px;

.orb1 {
  animation: spin 4s linear infinite;
  width: $size;
  height: $size;
  // margin-top: auto;
  // margin-bottom: auto;
  // margin-top: 140px;
  border-radius: 50%;
  box-shadow:
    inset 0 0 50px #fff,
    inset 20px 0 60px $violet1,
    inset -20px 0 60px $blue1,
    inset 20px 0 300px $violet1,
    inset -20px 0 300px $blue1,
    0 0 50px #fff,
    -10px 0 60px $violet1,
    10px 0 60px $blue1;
}

.orb2 {
  animation: spin 4s linear infinite;
  width: $size;
  height: $size;
  // margin-top: auto;
  // margin-bottom: auto;
  // margin-top: 140px;
  border-radius: 50%;
  box-shadow:
    inset 0 0 50px #fff,
    inset 20px 0 60px $violet2,
    inset -20px 0 60px $blue2,
    inset 20px 0 300px $violet2,
    inset -20px 0 300px $blue2,
    0 0 50px #fff,
    -10px 0 60px $violet2,
    10px 0 60px $blue2;
}

.orb3 {
  animation: spin 4s linear infinite;
  width: $size;
  height: $size;
  // margin-top: auto;
  // margin-bottom: auto;
  // margin-top: 140px;
  border-radius: 50%;
  box-shadow:
    inset 0 0 50px #fff,
    inset 20px 0 60px $violet3,
    inset -20px 0 60px $blue3,
    inset 20px 0 300px $violet3,
    inset -20px 0 300px $blue3,
    0 0 50px #fff,
    -10px 0 60px $violet3,
    10px 0 60px $blue3;
}

@keyframes spin {
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
</style>