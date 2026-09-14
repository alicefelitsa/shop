<template>
  <div class="content">
    <el-card shadow="always">
      <div slot="header" class="setting-header">
        <span>系统设置</span>
      </div>
      <div class="setting-form" v-loading="loading">
        <el-form ref="form" :model="form" label-width="90px">
          <el-form-item label="网站域名">
            <el-input v-model="form.domain" placeholder="如：http://127.0.0.1:8100"></el-input>
          </el-form-item>
          <el-form-item label="访问方式">
            <el-radio-group v-model="form.access_mode">
              <el-radio label="pc">PC</el-radio>
              <el-radio label="h5">H5</el-radio>
              <el-radio label="all">全部</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item class="setting-submit">
            <el-button type="primary" icon="el-icon-check" :loading="saving" @click="save">保存</el-button>
            <span class="setting-tip">域名用于前台产品图片等资源地址的拼接</span>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
  </div>
</template>

<script>
import {GetConfigSetting, SaveConfigSetting} from "@/api/config";

export default {
  name: "setting",
  data() {
    return {
      loading: false,
      saving: false,
      form: {
        domain: '',
        access_mode: 'all'
      }
    }
  },
  mounted() {
    this.getSetting()
  },
  methods: {
    //获取系统配置
    async getSetting() {
      this.loading = true
      try {
        let data = await GetConfigSetting()
        if (data && data.length > 0) {
          this.form = {
            domain: data[0].domain || '',
            access_mode: data[0].access_mode || 'all'
          }
        }
      } catch (e) {
        this.$message.error(e.message);
      } finally {
        this.loading = false
      }
    },
    //保存系统配置
    async save() {
      this.saving = true
      try {
        let message = await SaveConfigSetting({...this.form})
        this.$message.success(message)
      } catch (e) {
        this.$message.error(e.message);
      } finally {
        this.saving = false
      }
    }
  }
}
</script>

<style scoped>
/* 标题行与联系方式页一致：行高32px + 底部10px留白 + 字号14px */
.setting-header {
  margin-bottom: 10px;
}

.setting-header span {
  display: inline-block;
  line-height: 32px;
  font-size: 14px;
  font-weight: normal;
  color: #606266;
}

/* 表单定宽单列，避免输入框过宽拉伸 */
.setting-form {
  max-width: 560px;
  margin-top: 10px;
}

.setting-form >>> .el-form-item {
  margin-bottom: 20px;
}

/* 保存按钮与说明文字同行，紧凑收尾 */
.setting-submit {
  margin-top: 8px;
  margin-bottom: 0;
}

.setting-tip {
  margin-left: 12px;
  font-size: 12px;
  color: #909399;
}
</style>
