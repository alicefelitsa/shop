<template>
  <div id="app">
    <template v-if="ready && !blankActive">
      <SiteHeader />
      <main class="main-content">
        <transition name="page" mode="out-in">
          <router-view />
        </transition>
      </main>
      <SiteFooter />
    </template>
    <div v-else class="blank-page"></div>
  </div>
</template>

<script>
import SiteHeader from './components/SiteHeader.vue'
import SiteFooter from './components/SiteFooter.vue'
import {accessState} from './utils/access'

export default {
  name: 'App',
  components: {
    SiteHeader,
    SiteFooter
  },
  computed: {
    // 访问校验完成前整页不渲染，避免头尾先出现再切空白的闪烁
    ready() {
      return accessState.ready
    },
    // 访问受限且展示类型为空白时，头尾与路由视图均不渲染，URL 保持当前路径
    blankActive() {
      return accessState.blank
    }
  }
}
</script>

<style scoped>
#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  overflow-x: hidden;
  max-width: 100%;
}

.main-content {
  flex: 1;
  overflow-x: hidden;
  max-width: 100%;
  display: flex;
  flex-direction: column;
}

/* 空白展示占位：撑满视口且不透出任何内容 */
.blank-page {
  flex: 1;
  background: var(--bg-white);
}
</style>
