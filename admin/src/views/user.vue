<template>
  <div class="content">
    <el-card shadow="always">
      <div slot="header">

        <!--搜索栏-->
        <div class="queryForm">
          <el-form :inline="true" :model="where" class="query-form-inline" size="small">
            <el-form-item label="标题">
              <el-input v-model="where.title" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="地区">
              <el-input v-model="where.region" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="格式">
              <el-input v-model="where.format" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="年份">
              <el-input v-model="where.year" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="分类">
              <el-select v-model="where.category" placeholder="请输入" clearable class="queryElInput">
                <el-option label="喜剧" value="喜剧"></el-option>
                <el-option label="动作" value="动作"></el-option>
                <el-option label="爱情" value="爱情"></el-option>
                <el-option label="情色" value="情色"></el-option>
                <el-option label="科幻" value="科幻"></el-option>
                <el-option label="剧情" value="剧情"></el-option>
                <el-option label="恐怖" value="恐怖"></el-option>
                <el-option label="犯罪" value="犯罪"></el-option>
                <el-option label="惊悚" value="惊悚"></el-option>
                <el-option label="冒险" value="冒险"></el-option>
                <el-option label="悬疑" value="悬疑"></el-option>
                <el-option label="奇幻" value="奇幻"></el-option>
                <el-option label="动画" value="动画"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="search">查询</el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>

      <!--工具栏-->
      <div class="toolbar">
        <el-button type="primary" size="small" icon="el-icon-plus" @click="dialogVisible=true">添加</el-button>
        <el-button type="danger" size="small" icon="el-icon-delete" @click="del">删除</el-button>
        <el-button size="small" icon="el-icon-download" @click="handleExport">导出</el-button>
      </div>

      <!--数据表格-->
      <el-table class="tableData" :data="tableData" :highlight-selection-row="true" height="calc(100vh - 182px)"
                :border="true"
                v-loading="visitorLoading" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"></el-table-column>
        <el-table-column prop="id" label="ID" min-width="50px">
          <template v-slot="{row}">
            {{ row.id }}
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题">
          <template v-slot="{row}">
            {{ row.title }}
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类">
          <template v-slot="{row}">
            {{ row.category }}
          </template>
        </el-table-column>
        <el-table-column prop="director" label="导演">
          <template v-slot="{row}">
            {{ row.director }}
          </template>
        </el-table-column>
        <el-table-column prop="region" label="地区">
          <template v-slot="{row}">
            {{ row.region }}
          </template>
        </el-table-column>
        <el-table-column prop="year" label="年份" width="70px">
          <template v-slot="{row}">
            {{ row.year }}
          </template>
        </el-table-column>
        <el-table-column prop="episode" label="集数" width="70px">
          <template v-slot="{row}">
            {{ row.episode }}
          </template>
        </el-table-column>
        <el-table-column prop="format" label="格式" width="70px">
          <template v-slot="{row}">
            {{ row.format }}
          </template>
        </el-table-column>
        <el-table-column prop="play_url" label="播放路径" min-width="120px">
          <template v-slot="{row}">
            {{ row.play_url }}
          </template>
        </el-table-column>
        <el-table-column prop="ctime" label="时间" min-width="110px">
          <template v-slot="{row}">
            {{ row.ctime }}
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="100px">
          <template v-slot="{row}">
            <el-button size="mini" @click="edit(row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!--表格分页-->
      <div style="margin-top: 10px; text-align: center;" class="currentPage">
        <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="where.page"
            :page-sizes="pageSizes"
            :page-size="where.limit"
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalData">
        </el-pagination>
      </div>

    </el-card>

    <!--编辑对话框-->
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
import {AddMovie, DelMovie, GetMovie, SaveMovie} from "@/api/movie";
import {exportToExcel} from "@/utils/excel";

export default {
  data() {
    return {
      tableData: [],
      totalData: 0,
      pageSizes: [30, 50, 100, 200],
      multipleSelection: [],
      where: {
        title: '',
        category: '',
        region: '',
        director: '',
        year: '',
        format: '',
        page: 1,
        limit: 30,
      },
      loading: null,
      file_url: '',
      visitorLoading: false,
      dialogVisible: false,
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
  mounted() {
    this.where.limit = this.pageSizes[0]
    this.getMovie()
  },
  methods: {
    //获取影片
    async getMovie() {
      this.visitorLoading = true;
      setTimeout(async () => {
        try {
          let data = await GetMovie({...this.where})
          console.log(data);
          this.tableData = data.data;
          this.totalData = data.count
        } catch (e) {
          this.$message.error(e.message);
        } finally {
          this.visitorLoading = false;
        }
      }, 200)
    },
    //查询
    search() {
      this.where.page = 1
      this.getMovie()
    },
    //页数
    handleSizeChange(val) {
      this.where.limit = val
      console.log(`每页 ${this.where.limit} 条`);
      this.getMovie()
    },
    //页码
    handleCurrentChange(val) {
      this.where.page = val
      console.log(`当前页: ${this.where.page}`);
      this.getMovie()
    },
    //多选
    handleSelectionChange(val) {
      this.multipleSelection = []
      val.forEach(item => {
        this.multipleSelection.push(item.id);
      })
    },
    //获取修改数据
    edit(row) {
      this.form = {...row}
      this.dialogVisible = true;
      //console.log(row)
    },
    //添加或修改
    async save() {
      try {
        const addOrSave = this.form.id ? SaveMovie : AddMovie;
        let message = await addOrSave({...this.form})
        this.$message.success(message)
        this.handleClose()
        await this.getMovie()
      } catch (e) {
        this.$message.error(e.message);
      }
    },
    //删除
    async del() {
      if (this.multipleSelection.length === 0) {
        this.$message.warning("请选择要删除的数据");
        return
      }
      this.$confirm('即将删除，是否继续?').then(async _ => {
        let ids = this.multipleSelection.join(',')
        //console.log(ids)
        try {
          let message = await DelMovie(ids)
          this.$message.success(message)
          await this.getMovie()
        } catch (e) {
          this.$message.error(e.message);
        }
      }).catch(_ => {
      });
    },
    //关闭编辑对话框
    handleClose() {
      this.dialogVisible = false
      Object.keys(this.form).forEach(key => {
        this.form[key] = ''
      })
      if (this.$refs.form) {
        this.$refs.form.clearValidate()
      }
    },
    //导出表格
    async handleExport() {
      const headers = ['标题', '分类', '导演', '地区', '年份', '格式', '集数', '播放路径', '评分', '介绍']
      // 将 tableData 转换为导出格式
      const rows = this.tableData.map(item => [
        item.title || '',
        item.category || '',
        item.director || '',
        item.region || '',
        item.year || '',
        item.format || '',
        item.episode || '',
        item.play_url || '',
        item.score || '',
        item.plot || ''
      ])
      await exportToExcel(headers, rows, '影片信息表')
    }
  }
}
</script>

<style>

</style>