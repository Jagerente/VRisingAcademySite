<template>
  <header class="navbar navbar-expand-lg fixed-top"
    style="text-transform: uppercase; letter-spacing: 0.35rem; font-size:21px">
    <nav class="container-xxl navbar-dark">
      <a class="navbar-brand" href="https://vk.com/vrisingacademy" target="blank">
        <img src="@/assets/images/ui/logo.webp" alt="" class="d-inline-block align-text-top logo" />
      </a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="nav nav-pills" role="tablist">
          <li class="nav-item">
            <a @click="$router.push('/')" href="#tabInfo" class="nav-link" data-bs-toggle="tab">Information</a>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" href="#" role="button" aria-haspopup="true"
              aria-expanded="false">Catalogue</a>
            <div class="dropdown-menu">
              <a class="dropdown-item" @click="$router.push('/items')" href="#" data-bs-toggle="tab">Items</a>
              <a class="dropdown-item disabled" href="#" data-bs-toggle="tab">Spells</a>
              <a class="dropdown-item disabled" href="#" data-bs-toggle="tab">Bestiary</a>
            </div>
          </li>
          <li class="nav-item">
            <a href="https://vrising-map.com/" target="blank" class="nav-link">Map</a>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle disabled" data-bs-toggle="dropdown" href="#" role="button"
              aria-haspopup="true" aria-expanded="false">Guides</a>
            <div class="dropdown-menu">
              <a class="dropdown-item" href="#" data-bs-toggle="tab">Mechanics</a>
              <a class="dropdown-item" href="#" data-bs-toggle="tab">Bosses</a>
              <a class="dropdown-item" href="#" data-bs-toggle="tab">Resources</a>
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
export default {
  name: "my-navbar",
};
</script>

<style scoped>
.dropdown-item:hover,
.dropdown-item:focus,
.dropdown-item.active,
.dropdown-item:active,
.nav-pills .nav-link.active,
.nav-pills .show>.nav-link {
  background-color: #1c1b29;
  background: #1c1b29;
}

.dropdown-item:hover,
.dropdown-item:focus {
  color: white
}

header {
  /* height: 75px; */
  background-color: #1c1b29;
  box-shadow: 0px 5px 5px #11111d;
  padding: 5px;
}

.dropdown-menu {
  background: #1c1b29;
}

.dropdown-item {
  color: white;
}

.logo {
  height: 60px;
}

* {
  user-select: none;
}
</style>