<template>
  <div class="account-login-container">
    <!-- 背景效果 -->
    <div class="background-effects">
      <div class="gradient-bg"></div>
      <div class="floating-element floating-1"></div>
      <div class="floating-element floating-2"></div>
      <div class="floating-element floating-3"></div>
    </div>

    <!-- 登录主容器 -->
    <div class="login-main">
      <!-- 左侧品牌区域 -->
      <!--      <div class="brand-section">
              <div class="brand-logo">
                <i class="el-icon-user"></i>
              </div>
              <h1 class="brand-name">账号登录</h1>
              <p class="brand-subtitle">请使用您的账号登录系统</p>

              <div class="features">
                <div class="feature-item">
                  <i class="el-icon-check"></i>
                  <span>单点登录</span>
                </div>
                <div class="feature-item">
                  <i class="el-icon-check"></i>
                  <span>快速访问</span>
                </div>
                <div class="feature-item">
                  <i class="el-icon-check"></i>
                  <span>安全验证</span>
                </div>
              </div>
            </div>-->

      <!-- 右侧表单区域 -->
      <div class="form-section">
        <div class="form-header">
          <h2 style="color: #368a88;">欢迎回来</h2>
          <!--          <p>请输入您的账号继续</p>-->
        </div>

        <!-- 登录表单 -->
        <el-form
            ref="loginForm"
            :model="form"
            :rules="rules"
            label-width="0"
            class="account-form"
            @submit.native.prevent
            @keyup.enter.native="handleLogin"
        >
          <!-- 账号输入框 -->
          <el-form-item prop="account">
            <el-input
                v-model="form.account"
                placeholder="请输入账号"
                size="medium"
                clearable
                class="account-input"
            >
              <template slot="prepend">
                <span class="input-label">账号</span>
              </template>
            </el-input>
          </el-form-item>

          <!-- 密码输入框 -->
          <el-form-item prop="password">
            <el-input
                v-model="form.password"
                placeholder="请输入密码"
                size="medium"
                clearable
                class="password-input"
                type="password"
            >
              <template slot="prepend">
                <span class="input-label">密码</span>
              </template>
            </el-input>
          </el-form-item>

          <!-- 登录按钮 -->
          <el-form-item>
            <el-button
                type="primary"
                :loading="loading"
                @click="handleLogin"
                class="login-btn"
            >
              {{ loading ? '登录中...' : '立即登录' }}
            </el-button>
          </el-form-item>

          <!-- 其他登录方式 -->
          <!--          <div class="other-login">
                      <div class="divider">
                        <span>其他方式</span>
                      </div>

                      <div class="login-methods">
                        <el-button
                            circle
                            class="method-btn wechat"
                            @click="loginByWechat"
                        >
                          <i class="el-icon-chat-dot-round"></i>
                        </el-button>
                        <el-button
                            circle
                            class="method-btn qq"
                            @click="loginByQQ"
                        >
                          <i class="el-icon-s-promotion"></i>
                        </el-button>
                        <el-button
                            circle
                            class="method-btn phone"
                            @click="loginByPhone"
                        >
                          <i class="el-icon-phone-outline"></i>
                        </el-button>
                      </div>
                    </div>-->

          <!-- 底部链接 -->
          <!--          <div class="form-footer">
                      <el-link
                          type="info"
                          :underline="false"
                          @click="handleHelp"
                          class="help-link"
                      >
                        <i class="el-icon-question"></i>
                        需要帮助？
                      </el-link>
                      <el-link
                          type="info"
                          :underline="false"
                          @click="handleRegister"
                          class="register-link"
                      >
                        <i class="el-icon-user-solid"></i>
                        注册账号
                      </el-link>
                    </div>-->
        </el-form>

        <!-- 协议声明 -->
        <div class="agreement">
          <p>登录即表示您同意
            <el-link type="primary" :underline="false">用户协议</el-link>
            和
            <el-link type="primary" :underline="false">隐私政策</el-link>
          </p>
        </div>
      </div>
    </div>

    <!-- 底部信息 -->
    <div class="global-footer">
      <p>© 2026 企业服务系统 版本 1.0.0</p>
      <!--      <p>服务热线: 400-888-8888 | 技术支持: support@example.com</p>-->
    </div>

    <!-- 帮助弹窗 -->
    <el-dialog
        title="需要帮助？"
        :visible.sync="showHelpDialog"
        width="400px"
        center
    >
      <div class="help-content">
        <p>如果您遇到登录问题，请尝试以下方法：</p>
        <ul>
          <li>1. 确认您的账号是否正确</li>
          <li>2. 检查网络连接是否正常</li>
          <li>3. 联系管理员获取支持</li>
        </ul>
        <p>客服工作时间：周一至周五 9:00-18:00</p>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showHelpDialog = false">关闭</el-button>
        <el-button type="primary" @click="contactSupport">联系客服</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {login} from "@/api/login";

export default {
  name: 'Login',
  data() {
    return {
      // 表单数据
      form: {
        account: '',
        password: ''
      },

      // 验证规则
      rules: {
        account: [
          {required: true, message: '请输入账号', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'}
        ]
      },

      // 加载状态
      loading: false,

      // 帮助弹窗
      showHelpDialog: false
    }
  },

  created() {
    // 自动聚焦输入框
    /*this.$nextTick(() => {
      this.$refs.loginForm.$el.querySelector('input')?.focus()
    })*/
  },

  methods: {
    // 处理登录
    handleLogin() {
      this.$refs.loginForm.validate((valid) => {
        if (!valid) {
          console.log('表单验证失败：', valid)
          return false
        }
        this.loading = true
        login(this.form).then((msg) => {
          this.loading = false;
          this.$message.success(msg);
          this.$router.push('/')
        }).catch((e) => {
          this.loading = false;
          this.$message.error(e.message);
        });
      })
    },

    // 联系客服
    contactSupport() {
      this.showHelpDialog = false
      this.$message({
        message: '客服邮箱已复制到剪贴板',
        type: 'success'
      })
      // 模拟复制客服邮箱
      const email = 'support@example.com'
      navigator.clipboard?.writeText(email)
    }
  }
}
</script>

<style scoped>
.account-login-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
  background: #f5f7fa;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", Arial, sans-serif;
}

/* 背景效果 */
.background-effects {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 0;
}

.gradient-bg {
  position: absolute;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #1890ff 0%, #52c41a 100%);
  opacity: 0.1;
}

.floating-element {
  position: absolute;
  border-radius: 50%;
  opacity: 0.3;
  animation: float 20s infinite linear;
}

.floating-1 {
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, #1890ff, transparent 70%);
  top: 10%;
  right: 10%;
  animation-delay: 0s;
}

.floating-2 {
  width: 250px;
  height: 250px;
  background: linear-gradient(135deg, #52c41a, transparent 70%);
  bottom: 20%;
  left: 10%;
  animation-delay: 7s;
  animation-duration: 18s;
}

.floating-3 {
  width: 200px;
  height: 200px;
  background: linear-gradient(135deg, #13c2c2, transparent 70%);
  top: 60%;
  right: 20%;
  animation-delay: 14s;
  animation-duration: 22s;
}

@keyframes float {
  0%, 100% {
    transform: translate(0, 0) rotate(0deg);
  }
  25% {
    transform: translate(20px, 20px) rotate(90deg);
  }
  50% {
    transform: translate(-20px, 10px) rotate(180deg);
  }
  75% {
    transform: translate(10px, -20px) rotate(270deg);
  }
}

/* 登录主容器 */
.login-main {
  display: flex;
  max-width: 1000px;
  width: 40%;
  /*background: #f5f5f5;*/
  /*background: linear-gradient(135deg, #06420a 0%, #ff00dd 100%);*/
  border-radius: 20px;
  /*box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);*/
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5), /* 主阴影 */ 0 0 0 1px rgba(255, 255, 255, 0.1) inset, /* 内边框高光 */ 0 1px 2px rgba(255, 255, 255, 0.2) inset; /* 顶部高光 */
  overflow: hidden;
  z-index: 1;
  margin: 20px;
  animation: slideUp 0.8s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 左侧品牌区域 */
.brand-section {
  flex: 1;
  background: linear-gradient(135deg, #1890ff 0%, #52c41a 100%);
  color: white;
  padding: 60px 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.brand-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><circle cx="50" cy="50" r="1" fill="white" opacity="0.1"/></svg>');
  background-size: 20px 20px;
  opacity: 0.5;
}

.brand-logo {
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 30px;
  font-size: 40px;
  backdrop-filter: blur(5px);
  border: 2px solid rgba(255, 255, 255, 0.3);
}

.brand-name {
  font-size: 36px;
  font-weight: 600;
  margin-bottom: 10px;
  letter-spacing: 2px;
}

.brand-subtitle {
  font-size: 16px;
  opacity: 0.9;
  margin-bottom: 40px;
  letter-spacing: 1px;
}

.features {
  margin-top: 40px;
}

.feature-item {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  font-size: 16px;
  padding: 10px 15px;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
  cursor: pointer;
}

.feature-item:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateX(10px);
}

.feature-item i {
  margin-right: 15px;
  font-size: 20px;
  color: #52c41a;
  background: white;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 右侧表单区域 */
.form-section {
  flex: 1;
  padding: 60px 50px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.form-header {
  text-align: center;
  margin-bottom: 40px;
}

.form-header h2 {
  font-size: 32px;
  color: #333;
  margin-bottom: 10px;
  font-weight: 600;
}

.form-header p {
  color: #666;
  font-size: 16px;
}

/* 表单样式 */
.account-form {
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
}

.account-input >>> .el-input-group__prepend {
  background: linear-gradient(135deg, #1890ff 0%, #52c41a 100%);
  border: none;
  color: white;
  width: 40px;
  justify-content: center;
  font-weight: 600;
}

.account-input >>> .el-input__inner {
  border-left: none;
  padding-left: 15px;
  height: 50px;
  font-size: 16px;
  border-color: #dcdfe6;
  transition: all 0.3s ease;
}

.account-input >>> .el-input__inner:focus {
  border-color: #1890ff;
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}

.account-input >>> .el-input__prefix {
  display: flex;
  align-items: center;
  padding-left: 10px;
  color: #1890ff;
}

.password-input >>> .el-input-group__prepend {
  background: linear-gradient(135deg, #ea1d43 0%, #52c41a 100%);
  border: none;
  color: white;
  width: 40px;
  justify-content: center;
  font-weight: 600;
}

.password-input >>> .el-input__inner {
  border-left: none;
  padding-left: 15px;
  height: 50px;
  font-size: 16px;
  border-color: #dcdfe6;
  transition: all 0.3s ease;
}

.password-input >>> .el-input__inner:focus {
  border-color: #1890ff;
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}

.password-input >>> .el-input__prefix {
  display: flex;
  align-items: center;
  padding-left: 10px;
  color: #1890ff;
}

.login-btn {
  width: 100%;
  height: 50px;
  background: linear-gradient(135deg, #1890ff 0%, #52c41a 100%);
  border: none;
  color: white;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 2px;
  border-radius: 8px;
  transition: all 0.3s ease;
  margin-top: 20px;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(24, 144, 255, 0.3);
}

.login-btn:active {
  transform: translateY(0);
}

.login-btn.is-loading {
  opacity: 0.8;
}

/* 其他登录方式 */
.other-login {
  margin: 30px 0;
}

.divider {
  display: flex;
  align-items: center;
  margin: 20px 0;
  color: #999;
  font-size: 14px;
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  height: 1px;
  background: #e8e8e8;
}

.divider span {
  padding: 0 15px;
}

.login-methods {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.method-btn {
  width: 50px;
  height: 50px;
  border: none;
  font-size: 20px;
  color: white;
  transition: all 0.3s ease;
}

.method-btn.wechat {
  background: #07c160;
}

.method-btn.qq {
  background: #12b7f5;
}

.method-btn.phone {
  background: #1890ff;
}

.method-btn:hover {
  transform: translateY(-3px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
}

/* 底部链接 */
.form-footer {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #f0f0f0;
}

.help-link,
.register-link {
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 5px;
}

.help-link i {
  color: #1890ff;
}

.register-link i {
  color: #52c41a;
}

/* 协议声明 */
.agreement {
  text-align: center;
  margin-top: 30px;
  font-size: 12px;
  color: #999;
  line-height: 1.6;
}

.agreement p {
  margin: 0;
}

/* 全局页脚 */
.global-footer {
  margin-top: 30px;
  text-align: center;
  color: #666;
  font-size: 12px;
  line-height: 1.8;
  z-index: 1;
  padding: 20px;
}

.global-footer p {
  margin: 5px 0;
}

/* 帮助弹窗内容 */
.help-content {
  padding: 10px 0;
}

.help-content p {
  margin: 10px 0;
  color: #333;
}

.help-content ul {
  margin: 15px 0;
  padding-left: 20px;
  color: #666;
}

.help-content li {
  margin: 8px 0;
  line-height: 1.6;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-main {
    flex-direction: column;
    width: 95%;
  }

  .brand-section {
    padding: 40px 20px;
  }

  .form-section {
    padding: 40px 20px;
  }

  .brand-name {
    font-size: 28px;
  }

  .form-header h2 {
    font-size: 24px;
  }

  .floating-1,
  .floating-2,
  .floating-3 {
    opacity: 0.1;
  }
}

@media (max-width: 480px) {
  .login-main {
    border-radius: 10px;
  }

  .brand-section {
    padding: 30px 15px;
  }

  .form-section {
    padding: 30px 15px;
  }

  .form-footer {
    flex-direction: column;
    gap: 10px;
    align-items: center;
  }
}
</style>