<template>
  <div>
    <el-dialog :title="form.id?'修改影片':'添加影片'" :visible.sync="dialogVisible" width="45%"
               class="responsive-dialog" :close-on-click-modal="false" :before-close="handleClose" top="7vh">
      <el-form ref="form" :model="form" label-width="80px">
        <el-row>
          <el-col :span="24">
            <el-form-item label="标题">
              <el-input v-model="form.title" placeholder="请输入标题"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="分类">
              <el-input v-model="form.category" placeholder="请输入分类"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="导演">
              <el-input v-model="form.director" placeholder="请输入导演"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地区">
              <el-input v-model="form.region" placeholder="请输入地区"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="年份">
              <el-input v-model="form.year" placeholder="请输入年份"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="格式">
              <el-select v-model="form.format" placeholder="请选择格式">
                <el-option label="mp4" value="mp4"></el-option>
                <el-option label="m3u8" value="m3u8"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="集数">
              <el-input v-model="form.episode" placeholder="请输入集数"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="播放路径">
              <el-input v-model="form.play_url" placeholder="请输入路径"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="评分">
              <el-input v-model="form.score" placeholder="请输入评分"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-form-item label="介绍">
              <el-input v-model="form.plot" type="textarea" :rows="5" placeholder="请输入内容"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose" size="small">取 消</el-button>
        <el-button type="primary" @click="save" size="small">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {AddMovie, SaveMovie} from "@/api/movie";
// 添加或修改
export default {
  props: {
    //弹窗是否打开
    visible: {
      type: Boolean,
      default: false
    },
    //修改回显的数据
    editData: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      form: {
        id: '',
        title: '',
        category: '',
        region: '',
        director: '',
        year: '',
        format: '',
        play_url: '',
        episode: '',
        plot: '',
        score: '',
      }
    }
  },
  computed: {
    dialogVisible: {
      get() {
        return this.visible
      },
      set(val) {
        this.$emit('update:visible', val)
      }
    }
  },
  watch: {
    editData: {
      handler(data) {
        if (data && Object.keys(data).length > 0) {
          this.form = {...data}
        }
      }
    }
  },
  methods: {
    //添加或修改
    async save() {
      try {
        const addOrSave = this.form.id ? SaveMovie : AddMovie;
        let message = await addOrSave({...this.form})
        this.$message.success(message)
        this.handleClose()
        this.$emit('done')
      } catch (e) {
        this.$message.error(e.message);
      }
    },
    //关闭编辑对话框
    handleClose() {
      this.dialogVisible = false
      this.resetForm()
    },
    //重置表单
    resetForm() {
      Object.keys(this.form).forEach(key => {
        this.form[key] = ''
      })
      if (this.$refs.form) {
        this.$refs.form.clearValidate()
      }
    }
  }
}
</script>

<style scoped>

</style>