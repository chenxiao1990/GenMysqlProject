<template>
  <div class="home">
    <div style="width:70%;float:left;">
      <div style="margin:20px;">
        <span>版本:{{version}}</span>

      </div>

      <el-form style="width:400px;margin: auto;" label-position="left" label-width="100px">
        <el-form-item label="数据库地址">
          <el-input v-model="dbIPPort" placeholder="192.168.0.24:3306"></el-input>
        </el-form-item>
        <el-form-item label="数据库名">
          <el-input v-model="dbName" placeholder="xzy2"></el-input>
        </el-form-item>
        <el-form-item label="数据库用户">
          <el-input v-model="dbUser" placeholder="root"></el-input>
        </el-form-item>
        <el-form-item label="数据库密码">
          <el-input v-model="dbPass" placeholder="12345678"></el-input>
        </el-form-item>
      </el-form>
      <div style="margin:20px;">
        <el-button @click="jsontogo" type="primary">json转go结构体</el-button>
      </div>
      <div style="margin:20px;">
        <el-button @click="gentable" type="primary">表生成struct</el-button>
      </div>
      <div style="margin:20px;">
        <el-button @click="genselectcode" type="primary">查询代码生成</el-button>
      </div>
      <div style="margin:20px;">
        <el-button @click="genproject" type="success">生成工程</el-button>
      </div>
    </div>

    <el-dialog title="自动生成工程" :visible.sync="dialogVisible" width="50%">
      新工程名<el-input v-model="outProjectName" placeholder="Mypro"></el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="genprojecting">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "home",
  components: {},
  data() {
    return {
      version: "",
      dbIPPort: "",
      dbName: "",
      dbUser: "",
      dbPass: "",
      outProjectName: "",
      dialogVisible: false
    };
  },
  created() {
    this.$axios.get("/cx/version").then(res => {
      console.log(res.data);
      this.version = res.data.data;
    });
    this.$axios.get("/cx/dbinfo").then(res => {
      console.log(res.data);
      this.dbIPPort = res.data.data.dbIPPort;
      this.dbName = res.data.data.dbName;
      this.dbUser = res.data.data.dbUser;
      this.dbPass = res.data.data.dbPass;
      this.outProjectName = res.data.data.outProjectName;
    });
  },
  methods: {
    gentable() {
      this.$axios
        .post("/cx/setdbinfo", {
          dbIPPort: this.dbIPPort,
          dbName: this.dbName,
          dbUser: this.dbUser,
          dbPass: this.dbPass
        })
        .then(res => {
          if (res.data.code == 1) {
            this.$router.push("/gentablestruct");
          } else {
            this.$message.error(res.data.message);
          }
        });
    },
    genproject() {
      this.dialogVisible = true;
    },
    genprojecting() {
      this.dialogVisible = false;
      this.$axios
        .post("/cx/genproject", {
          dbIPPort: this.dbIPPort,
          dbName: this.dbName,
          dbUser: this.dbUser,
          dbPass: this.dbPass,
          outProjectName: this.outProjectName
        })
        .then(res => {
          if (res.data.code == 1) {
            this.$message.success(res.data.message);
          } else {
            this.$message.error(res.data.message);
          }
        });
    },
    genselectcode() {
      this.$axios
        .post("/cx/setdbinfo", {
          dbIPPort: this.dbIPPort,
          dbName: this.dbName,
          dbUser: this.dbUser,
          dbPass: this.dbPass
        })
        .then(res => {
          if (res.data.code == 1) {
            this.$router.push("/genselectcode");
          } else {
            this.$message.error(res.data.message);
          }
        });
    },
    jsontogo() {
      this.$router.push("/jsontogo");
    }
  }
};
</script>
