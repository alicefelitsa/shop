<template>
  <el-dialog
      title="二维码"
      :visible.sync="dialogVisible"
      :width="screenWidth < 768 ? '100%' : '600px'"
      :close-on-click-modal="false"
      @close="updateVisible"
      top="10vh"
  >

    <el-tabs v-model="activeName" type="border-card" style="height: 390px;">
      <el-tab-pane label="QQ" name="qq">
        <div class="qrcode-dialog-content">
          <!-- 显示二维码 -->
          <div class="qrcode-section">
            <div class="qrcode-preview">
              <!-- 使用 img 显示二维码 -->
              <img
                  :src="qqQrcode"
                  alt="二维码"
                  class="qrcode-image"
                  v-if="qqQrcode"
              />
              <!-- 二维码信息 -->
              <div class="qrcode-info">
                <p class="info-content">{{ qqLink }}</p>
              </div>
            </div>
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane label="微信" name="wx">
        <div class="qrcode-dialog-content">
          <!-- 显示二维码 -->
          <div class="qrcode-section">
            <div class="qrcode-preview">
              <!-- 使用 img 显示二维码 -->
              <img
                  :src="wxQrcode"
                  alt="二维码"
                  class="qrcode-image"
                  v-if="wxQrcode"
              />
              <!-- 二维码信息 -->
              <div class="qrcode-info">
                <p class="info-content">{{ wxLink }}</p>
              </div>
            </div>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>
    <div class="button-container">
      <el-button style="margin: 20px 0 0 0;" type="primary" @click="getQr">重新获取</el-button>
    </div>
  </el-dialog>
</template>

<script>
import QRCode from "qrcode";
import {getQrcodeUrl} from "@/api/qr";

export default {
  props: {
    visible: Boolean
  },
  data() {
    return {
      activeName: 'qq',
      dialogVisible: false,
      screenWidth: 0,
      generating: false,        // 生成中状态
      qrcodeOptions: {
        errorCorrectionLevel: 'H',  // 容错级别: L/M/Q/H
        margin: 1,                   // 边距
        width: 240,                  // 宽度
        color: {
          dark: '#000000',          // 前景色
          light: '#FFFFFF'          // 背景色
        }
      },
      qqQrcode: '',            // QQ二维码图片URL
      qqLink: '',     // QQ链接
      wxLink: '',
      wxQrcode: ''
    }
  },
  watch: {
    visible(status) {
      if (status) {
        this.dialogVisible = true;
        this.getQr()
      } else {
        this.dialogVisible = false;
        console.log('二维码窗口关闭')
      }
    }
  },
  mounted() {
    // 初始化屏幕宽度
    this.screenWidth = window.innerWidth
    // 添加 resize 事件监听
    window.addEventListener('resize', this.handleResize)
  },
  beforeDestroy() {
    // 移除事件监听，防止内存泄漏
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
    //获取当前屏幕宽度
    handleResize() {
      this.screenWidth = window.innerWidth
      console.log('当前屏幕宽度:', this.screenWidth)
    },
    //提取二维码
    getQr() {
      getQrcodeUrl(localStorage.getItem("ucode")).then(async (data) => {
        this.$message.success('获取完成')
        if (data.qqLink) {
          this.qqLink = data.qqLink;
          this.qqQrcode = await this.generateQRCode(this.qqLink)
        } else {
          this.qqQrcode = ''
          this.qqLink = '无可用链接'
        }
        if (data.wxLink) {
          this.wxLink = data.wxLink;
          this.wxQrcode = await this.generateQRCode(this.wxLink)
        } else {
          this.wxQrcode = ''
          this.wxLink = '无可用链接'
        }
      }).catch((e) => {
        this.$message.error(e.message);
      });

    },
    // 生成二维码
    async generateQRCode(text) {
      if (!text.trim()) {
        this.$message.warning('请输入要生成二维码的内容')
        return false
      }
      this.generating = true
      try {
        //this.$message.success('二维码生成成功！')
        return await QRCode.toDataURL(text, this.qrcodeOptions)
      } catch (error) {
        this.$message.error('二维码生成失败，请重试')
        return false
      } finally {
        this.generating = false
      }
    },
    /* 更新visible */
    updateVisible(value) {
      this.activeName = 'qq'
      this.$emit('update:visible', value);
    }
  }
}
</script>

<style scoped>
/* 二维码预览区域：适配弹窗宽度 */
.qrcode-dialog-content {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  overflow-x: hidden;
}

/* 二维码图片/Canvas：自适应容器 */
.qrcode-image,
.qrcode-canvas {
  max-width: 100%; /* 防止图片/Canvas 超出弹窗宽度 */
  height: auto; /* 保持宽高比 */
}

.info-content {
  font-size: 16px;
  margin: 30px 0 20px 0;
}

.button-container {
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中（如果需要） */
  width: 100%;
}
</style>