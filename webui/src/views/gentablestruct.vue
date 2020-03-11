<template>
  <div>
    <div class="lefttable">
      <div style="border-bottom:1px solid #aaa;background-color:#409EFF; color:#FFF;line-height:30px;">
        <span>数据库表</span>
      </div>
      <div v-for=" (t,i) in tables" :key="i" class="tablecell">
        <div style="width:calc(100% - 40px);float:left">
          {{t.TableName}}
        </div>
        <div style="width: 38px ;float:right;">

          <el-button @click="selecttable(t)" type="success" size="mini" icon="el-icon-check" circle></el-button>
        </div>
        <div style="clear:both;"></div>
      </div>
    </div>
    <div class="center">
      <el-input type="textarea" style="color: #000;" rows="25" placeholder="请输入内容" v-model="gocode" show-word-limit></el-input>
    </div>

  </div>
</template>


<script>
export default {
  name: "genselectcode",
  components: {},
  computed: {},
  data() {
    return {
      tables: [], //所有数据库表
      maptables: {}, // 表名和表的映射

      curTables: [], //当前选择了的数据表
      gocode: ""
    };
  },
  created() {
    //获取所有数据库表信息
    this.inittables();
  },
  methods: {
    inittables() {
      this.$axios.get("/cx/tables").then(res => {
        if (res.data.code == 0) {
          this.$message.error(res.data.message);
        } else {
          this.tables = res.data.data;
          this.tables.forEach(e => {
            this.maptables[e.TableName] = e;
          });
          console.log("tables:", this.tables);
        }
      });
    },

    selecttable(t) {
      console.log(t);
      this.curTables = [];
      this.curTables.push(t);
      this.$axios.post("/cx/gentablestruct", t).then(res => {
        if (res.data.code == 1) {
          this.gocode = res.data.data;
        } else {
          this.gocde = res.data.message;
        }
      });
    }
  }
};
</script>

<style >
.lefttable {
  width: 300px;
  float: left;
  height: 100vh;
  border-right: 1px solid #aaa;
  overflow: auto;
}
.center {
  width: calc(100vw - 402px);
  float: left;
  height: 100vh;
  overflow: auto;
}
.right {
  width: 100px;
  float: left;
  height: 100vh;
  border-left: 1px solid #aaa;
  overflow: auto;
}
.tablecell {
  padding-left: 5px;
  text-align: left;
  border-bottom: 1px solid #ddd;
  line-height: 30px;
}

.onetable {
  width: 200px;
  height: 300px;
  margin: 20px;
  border: 1px solid #409eff;
  border-radius: 4px;
  float: left;
}
.gongnengbtn {
  margin-top: 20px !important;
}
</style>