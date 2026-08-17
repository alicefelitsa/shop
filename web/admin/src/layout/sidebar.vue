<template>
  <div class="sidebar-container">
    <div class="avatar-section">
      <el-popover
          placement="right-start"
          width="220"
          trigger="hover"
      >
        <div class="userInfo">昵称：{{ account ?? '在线客服' }}</div>
        <div class="userInfo">账号：{{ account }}</div>
        <div class="userInfo">到期时间：永久</div>
        <el-avatar slot="reference" :size="55" :src="avatar"/>
      </el-popover>

    </div>
    <el-menu :default-active="$route.path" router style="border:0" class="menu">
      <el-menu-item
          v-for="item in menuList"
          :key="item.path"
          :index="item.path"
          @click="item.path ? null : dialogVisible = true"
          class="vertical-menu-item"
      >
        <div class="menu-item-content">
          <i :class="item.icon" class="menu-icon"></i>
          <span class="menu-title">{{ item.title }}</span>
        </div>
      </el-menu-item>

      <el-menu-item
          class="vertical-menu-item logout-item"
          @click="handleLogout"
      >
        <div class="menu-item-content">
          <i class="el-icon-switch-button menu-icon"></i>
          <span class="menu-title">退出登录</span>
        </div>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
import {EventBus} from "@/utils/event-bus";

export default {
  data() {
    return {
      avatar: '/static/avatar.jpg',
      account: localStorage.getItem("account"),
      menuList: [
        {path: '/product', title: '产品', icon: 'el-icon-goods'},
        {path: '/category', title: '分类', icon: 'el-icon-menu'},
        {path: '/message', title: '留言', icon: 'el-icon-chat-dot-square'},
        {path: '/contact', title: '联系方式', icon: 'el-icon-phone-outline'},
        {path: '/setting', title: '设置', icon: 'el-icon-setting'},
      ],
      dialogVisible: false,
    }
  },
  created() {
    EventBus.$on('sidebar', this.fetchData)
  },
  beforeDestroy() {
    EventBus.$off('sidebar', this.fetchData)
  },
  methods: {
    //重置data中的所有数据
    fetchData(status) {
      if (status) {
        Object.assign(this.$data, this.$options.data.call(this))
      }
    },
    //退出登录
    handleLogout() {
      this.$confirm('确定要退出登录吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 执行退出登录逻辑
        localStorage.removeItem('token')
        localStorage.removeItem('account')
        this.$router.push('/login')
      }).catch(() => {
        // 取消退出
      })
    },
  }
}
</script>

<style scoped>
.vertical-menu-item {
  height: auto !important; /* 允许自动高度 */
  line-height: normal !important; /* 重置行高 */
  padding: 8px 20px !important; /* 调整内边距 */
  margin: 0 0 4px 0 !important; /* 上下8px，左右0 */
}

.vertical-menu-item .menu-item-content {
  display: flex;
  flex-direction: column; /* 垂直排列 */
  align-items: center; /* 水平居中 */
  justify-content: center; /* 垂直居中 */
  text-align: center; /* 文字居中 */
}

.vertical-menu-item .menu-icon {
  font-size: 26px; /* 图标大小 */
  margin-bottom: 4px; /* 图标和文字的间距 */
  display: block; /* 块级显示 */
}

.vertical-menu-item .menu-title {
  font-size: 12px; /* 文字大小 */
  line-height: 1.2; /* 行高 */
  white-space: normal; /* 允许换行 */
  word-break: break-all; /* 长单词换行 */
}

/* 鼠标悬停效果 */
.logout-item:hover {
  background-color: #f56c6c !important; /* 红色背景 */
  color: #fff !important;
}

.logout-item:hover .menu-icon,
.logout-item:hover .menu-title {
  color: #fff !important;
}

.sidebar-container {
  display: flex;
  flex-direction: column; /* 垂直排列（头像在上，菜单在下） */
  align-items: center; /* 水平居中（头像和菜单都居中） */
  justify-content: flex-start; /* 子元素从顶部开始排列（可根据需求改为 center / space-between） */
  height: 100vh; /* 占满整个屏幕高度（或继承父容器高度，如 height: 100%） */
  padding: 10px 0 0 0; /* 上下内边距，避免内容贴边 */
  background-color: #fff; /* 侧边栏背景（可选） */
  box-sizing: border-box; /* 确保padding不影响高度 */
}

.avatar-section {
  margin-bottom: 10px; /* 头像与菜单之间的间距 */
}

.menu {
  width: 100%; /* 菜单宽度占满侧边栏 */
  flex: 1; /* 占满剩余垂直空间（使菜单在头像下方拉伸） */
  overflow-y: auto; /* 如果菜单项过多，允许垂直滚动 */
}

/* 菜单项样式（可选，优化美观） */
.vertical-menu-item {
  display: flex;
  justify-content: center; /* 菜单项内容水平居中 */
}

.menu-item-content {
  display: flex;
  align-items: center;
  gap: 4px; /* 图标和文字的间距 */
}

.userInfo {
  margin-bottom: 5px;
}
</style>