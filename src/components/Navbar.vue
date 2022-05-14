<template>
  <header class="navbar navbar-expand-md sticky-top">
    <nav class="container-xxl navbar-dark">
      <a class="navbar-brand" href="https://vk.com/vrising">
        <img
          src="@/assets/images/ui/logo.png"
          alt=""
          class="d-inline-block align-text-top logo"
        />
      </a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="nav nav-pills" role="tablist">
          <li class="nav-item">
            <a
              @click="$router.push('/')"
              href="#tabInfo"
              class="nav-link"
              data-bs-toggle="tab"
              >Information</a
            >
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              data-bs-toggle="dropdown"
              href="#"
              role="button"
              aria-haspopup="true"
              aria-expanded="false"
              >Catalogue</a
            >
            <div class="dropdown-menu">
              <a
                class="dropdown-item"
                @click="$router.push('/items')"
                href="#"
                data-bs-toggle="tab"
                >Items</a
              >
              <a class="dropdown-item disabled" href="#" data-bs-toggle="tab"
                >Spells</a
              >
              <a class="dropdown-item disabled" href="#" data-bs-toggle="tab"
                >Bestiary</a
              >
            </div>
          </li>
          <li class="nav-item">
            <a
              href="#"
              @click="$router.push('/map')"
              class="nav-link disabled"
              data-bs-toggle="tab"
              >Map</a
            >
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle disabled"
              data-bs-toggle="dropdown"
              href="#"
              role="button"
              aria-haspopup="true"
              aria-expanded="false"
              >Guides</a
            >
            <div class="dropdown-menu">
              <a class="dropdown-item" href="#" data-bs-toggle="tab"
                >Mechanics</a
              >
              <a class="dropdown-item" href="#" data-bs-toggle="tab">Bosses</a>
              <a class="dropdown-item" href="#" data-bs-toggle="tab"
                >Resources</a
              >
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
</script>

<style scoped>
header {
  /* height: 75px; */
  background-color: #1c1b29;
  box-shadow: 0px 5px 5px #11111d;
  padding: 5px;
}

/* .nav-link.active {
  background: center/100% no-repeat
    url("https://psv4.userapi.com/c237231/u205977252/docs/d48/9bf5c741893a/krug1.png?extra=s3KQPldmOdPwOukHqDfiefE1k2ddywwnj3Ty9GHfBBk-0jEWX56vGe5qdfgLSicqCxv88OSAMhkl4CZIV2-hecAgBmvmj2R5khhNvzXjlyvZlzJ_MyUd4da8hPn66ndQr90JpBXzroo05maPqidlsDyt");
} */

.dropdown-menu {
  background: #1c1b29;
}

.dropdown-item {
  color: white;
}

.logo {
  height: 45px;
}

* {
  user-select: none;
}
</style>