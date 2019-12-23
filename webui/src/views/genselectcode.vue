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

      <div class="onetable" v-for="(t,i) in curTables" :key="t.TableName">
        <div style=" border-bottom:1px solid #ddd;text-align:left;">
          <el-button @click="quanxuan(t)" type="primary" size="mini" circle>全</el-button>
          <span style="line-height: 30px;"> {{t.TableName}} </span>

        </div>
        <!-- 列出字段 -->
        <div style="height:calc(100% - 60px);overflow: auto; text-align: left;padding-left: 5px;">
          <div v-for="f in t.Fields" :key="f.Name">
            <div style="line-height:25px;">
              <el-checkbox @change="checkchange()" v-model="f.checked">{{f.Name}} &nbsp;&nbsp;&nbsp;&nbsp;{{f.Type}}</el-checkbox>
              <div v-if="guanlian[t.TableName] != null && guanlian[t.TableName].guanlianField == f.Name" style="color:#f00;">
                {{f.Name}} 关联 {{guanlian[t.TableName].guanlian2Table}}.{{guanlian[t.TableName].guanlian2Field}}
              </div>
            </div>

          </div>
        </div>

        <div style="height:30px;">
          <el-button @click="setzhu(i)" type="danger" size="mini" circle>{{i == 0 ? '主': '副'}}</el-button>
          <el-button v-if="i != 0" @click="setlianjie(i)" type="danger" size="mini" circle>联</el-button>
          <el-button @click="setmore(i)" v-if="i != 0" type="danger" size="mini" circle>{{t.ismore == 0 ? '1对1': '1对多'}}</el-button>
          <el-button @click="removetable(t)" type="danger" size="mini" icon="el-icon-close" circle></el-button>
        </div>

      </div>

    </div>
    <div class="right">
      <el-button @click="clear()" class="gongnengbtn" type="danger" size="mini">清空</el-button>

      <el-button @click="addtiaojian()" class="gongnengbtn" type="primary" size="mini">编辑条件</el-button>
      <el-button @click="gencode()" class="gongnengbtn" type="success" size="mini">生成代码</el-button>
    </div>
    <div style="clear:both;"> </div>

    <el-dialog title="关联表" style="text-align:left;" v-if="dialogVisible" :visible.sync="dialogVisible" width="50%">

      <div>
        本表字段:<el-select v-model="guanlianField" placeholder="请选择本表字段">
          <el-option v-for="item in maptables[guanlianTable].Fields" :key="item.Name" :label="guanlianTable+'.'+item.Name" :value="item.Name">
          </el-option>
        </el-select>
      </div>
      <div>
        关联表:<el-select v-model="guanlian2Table" placeholder="请选择本表字段">
          <el-option v-for="item in canselecttables" :key="item" :label="item" :value="item">
          </el-option>
        </el-select>
      </div>

      <div v-if="guanlian2Table != ''">
        关联字段:<el-select v-model="guanlian2Field" placeholder="请选择关联表字段">
          <el-option v-for="item in maptables[guanlian2Table].Fields" :key="item.Name" :label="item.Name" :value="item.Name">
          </el-option>
        </el-select>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="guanliansave">添加关联</el-button>

      </span>
    </el-dialog>

    <el-dialog :title="'编辑条件('+ tiaojianSelect.length+')'" style="text-align:left;" v-if="tiaojiandialog" :visible.sync="tiaojiandialog" width="50%">

      <div>
        <div v-for="(tx,i) in tiaojianSelect" :key="i">
          {{tx}} <el-button type="danger" @click="deletetiaojian(i)" circle icon="el-icon-delete"></el-button>
        </div>
        <el-select style="width:100px;" v-model="tiaojiantmp" placeholder="符号">
          <el-option v-for="item in tiaojians" :key="item" :label="item" :value="item">
          </el-option>
        </el-select>
        <el-select v-model="tiaojianfiledtmp" placeholder="字段">
          <el-option v-for="item in tiaojianfieldsList" :key="item.Name" :label="item.Name" :value="item.Name">
          </el-option>
        </el-select>
        <el-button type="primary" @click="addtiaojian2" circle icon="el-icon-circle-plus-outline"></el-button>
      </div>

    </el-dialog>

    <el-dialog title="代码" style="text-align:left;" top="5%" v-if="codedialog" :visible.sync="codedialog" width="90%">
      <div style="width:33%; float:left;">
        {{curTables[0].TableName}}Api.go
        <el-input type="textarea" rows="20" placeholder="请输入内容" v-model="apicode" show-word-limit></el-input>
      </div>
      <div style="width:33%; float:left;">
        {{curTables[0].TableName}}Service.go
        <el-input type="textarea" rows="20" placeholder="请输入内容" v-model="servicecode" show-word-limit></el-input>
      </div>
      <div style="width:33%; float:left;">
        {{curTables[0].TableName}}Dao.go
        <el-input type="textarea" rows="20" placeholder="请输入内容" v-model="daocode" show-word-limit></el-input>
      </div>
      <div style="clear:both;"></div>
    </el-dialog>
  </div>
</template>


<script>
export default {
  name: "genselectcode",
  components: {},
  computed: {
    canselecttables() {
      let bb = [];
      this.curTables.forEach(element => {
        if (element.TableName != this.guanlianTable) {
          bb.push(element.TableName);
        }
      });
      return bb;
    }
  },
  data() {
    return {
      tables: [], //所有数据库表
      maptables: {}, // 表名和表的映射

      curTables: [], //当前选择了的数据表

      guanlianField: "",
      guanlianTable: "",
      guanlian2Field: "",
      guanlian2Table: "",
      dialogVisible: false, //关联对话框

      guanlian: {}, //关联记录

      tiaojiandialog: false, //添加条件
      tiaojians: [
        "=",
        ">",
        ">=",
        "<",
        "<=",
        "in",
        "like",
        "升序",
        "降序",
        "pageNum",
        "pageSize"
      ],
      tiaojianfieldsList: [], //所有条件字段
      tiaojianSelect: [], // 选定的条件  [拼接的字符串]
      tiaojianOri: {}, // 拼接的字符串:{OrName:,OrType:} 记录参数名和类型
      tiaojiantmp: "",
      tiaojianfiledtmp: "",

      //生成de代码
      daocode: "",
      servicecode: "",
      apicode: "",

      codedialog: false //代码窗口
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
      let bhave = false;
      this.curTables.forEach(tmp => {
        if (tmp.TableName == t.TableName) {
          bhave = true;
        }
      });
      if (bhave) {
        return;
      }

      t.ismore = 0;
      t.Fields.forEach(e => {
        e.checked = false;
      });

      this.curTables.push(t);
    },
    removetable(t) {
      var index = -1;
      this.curTables.forEach((tmp, i) => {
        if (tmp.TableName == t.TableName) {
          index = i;
        }
      });
      if (index > -1) {
        this.curTables.splice(index, 1);
        delete this.guanlian[t.TableName];
      }
    },
    setzhu(i) {
      let tmp = this.curTables[i];
      tmp.ismore = 0;
      this.curTables.splice(i, 1);
      this.curTables.unshift(tmp);
    },
    setmore(i) {
      this.curTables[i].ismore = this.curTables[i].ismore == 0 ? 1 : 0;
      this.curTables = JSON.parse(JSON.stringify(this.curTables));
    },
    clear() {
      this.curTables = [];
      this.guanlian = {};
      this.tiaojianSelect = [];
      this.tiaojianOri = {};
    },
    quanxuan(t) {
      t.Fields.forEach(e => {
        e.checked = e.checked == false ? true : false;
      });

      this.curTables = JSON.parse(JSON.stringify(this.curTables));
    },
    checkchange() {
      this.curTables = JSON.parse(JSON.stringify(this.curTables));
    },
    setlianjie(i) {
      this.guanlianTable = this.curTables[i].TableName;
      console.log(
        "当前操作关联表:",
        this.guanlianTable,
        this.maptables[this.guanlianTable]
      );
      this.dialogVisible = true;
    },
    guanliansave() {
      this.guanlian[this.guanlianTable] = {
        guanlianField: this.guanlianField,
        guanlian2Field: this.guanlian2Field,
        guanlian2Table: this.guanlian2Table
      };
      this.dialogVisible = false;
      this.guanlianField = "";
      this.guanlianTable = "";
      this.guanlian2Field = "";
      this.guanlian2Table = "";
    },
    deletetiaojian(i) {
      delete this.tiaojianOri[this.tiaojianSelect[i]];
      this.tiaojianSelect.splice(i, 1);
    },
    addtiaojian() {
      this.tiaojianfieldsList = [];
      this.curTables.forEach(element => {
        element.Fields.forEach(f => {
          let tmp = {
            OrName: f.Name,
            Name: element.TableName + "." + f.Name,
            Type: f.Type
          };
          this.tiaojianfieldsList.push(tmp);
        });
      });
      this.tiaojiandialog = true;
    },
    addtiaojian2() {
      if (this.tiaojiantmp == "") {
        this.$message.error("先选择条件内容");
        return;
      }
      let tmpsplit1 = this.tiaojianfiledtmp.split(".");
      let tiaojiankey = "";
      for (let i = 0; i < this.curTables.length; i++) {
        let element = this.curTables[i];
        if (element.TableName == tmpsplit1[0]) {
          tiaojiankey = "t" + i + "." + tmpsplit1[1];
          break;
        }
      }
      if (this.tiaojiantmp == "降序") {
        this.tiaojianSelect.push('Order("' + tiaojiankey + ' desc").\n');
      } else if (this.tiaojiantmp == "升序") {
        this.tiaojianSelect.push('Order("' + tiaojiankey + ' asc").\n');
      } else if (this.tiaojiantmp == "pageNum") {
        let comt = "Offset((pageNum - 1)*pageSize).\n";
        this.tiaojianSelect.push(comt);
        this.tiaojianOri[comt] = {
          OrName: "pageNum",
          OrType: "int"
        };
      } else if (this.tiaojiantmp == "pageSize") {
        let comt = "Limit(pageSize).\n";
        this.tiaojianSelect.push(comt);
        this.tiaojianOri[comt] = {
          OrName: "pageSize",
          OrType: "int"
        };
      } else {
        let orname = "";
        let ortype = "";
        this.tiaojianfieldsList.forEach(ef => {
          if (ef.Name == this.tiaojianfiledtmp) {
            orname = ef.OrName;
            ortype = this.sqltypetogotype(ef.Type);
          }
        });
        let tmpllll = " ? ";
        if (this.tiaojiantmp == "in") {
          tmpllll = " (?) ";
          orname += "s";
          ortype = "[]" + ortype;
        }

        let tmporname = orname;
        if (this.tiaojiantmp == "like") {
          tmporname = '"%"+' + orname + '+"%"';
        }
        let comt =
          'Where("' +
          tiaojiankey +
          " " +
          this.tiaojiantmp +
          tmpllll +
          '", ' +
          tmporname +
          ").\n";
        this.tiaojianSelect.push(comt);

        this.tiaojianOri[comt] = {
          OrName: orname,
          OrType: ortype
        };
      }
      this.tiaojiantmp = "";
      this.tiaojianfiledtmp = "";
    },
    sqltypetogotype(sqltype) {
      switch (sqltype.toLowerCase()) {
        case "tinyint":
        case "int":
        case "smallint":
        case "mediumint":
          return "int";
        case "bigint":
          return "int64";
        case "char":
        case "enum":
        case "varchar":
        case "longtext":
        case "mediumtext":
        case "text":
        case "tinytext":
          return "string";
        case "date":
        case "datetime":
        case "time":
        case "timestamp":
          return "time.Time";
        case "decimal":
        case "double":
          return "float64";
        case "float":
          return "float";
        case "binary":
        case "blob":
        case "longblob":
        case "mediumblob":
        case "varbinary":
          return "[]byte";
      }
      return "";
    },
    // 返回node 内的字段
    dfs(node, tableMore) {
      let refields = [];
      //node  {TableName:"", Fields:[{key:"",type:"",tag:""}],links:[]}
      if (node.links.length > 0) {
        node.links.forEach(element => {
          let mml = this.dfs(element, tableMore);
          mml.forEach(xx => {
            refields.push(xx);
          });
        });
        node.links = [];
        refields.forEach(ex => {
          node.Fields.push(ex);
        });
        return this.dfs(node, tableMore);
      } else {
        if (tableMore[node.TableName] == 1) {
          //对多
          let structstr = "[]struct{ \n";
          node.Fields.forEach(filed => {
            structstr += filed.key + " " + filed.type + " " + filed.tag + "\n";
          });
          structstr += "}";
          let mml = {
            key: "T" + node.TableName,
            type: structstr,
            tag: `\"json:\\\"` + node.TableName + `\\\"\"`
          };
          refields.push(mml);
          return refields;
        } else {
          //对单  直接合并到父级

          node.Fields.forEach(ee => {
            refields.push(ee);
          });
          return refields;
        }
      }
      return refields;
    },
    //这里是重头戏 自动生成代码
    gencode() {
      if (this.curTables.length == 0) {
        this.$message.error("先添加数据表,并勾选查询字段");
        return;
      }
      // 收集table名称  [tablename]
      let tablenames = [];
      // table重命名  tablename: rename
      let tableRename = {};
      // table ->字段 表中勾选的字段  tablename: Fields
      let tableFields = {};
      // 字段别名  table.filed: bieming
      let fieldsRename = {};
      // table ->more  表关联是否是1对多 tablename: 0/1
      let tableMore = {};
      //关联信息 table->info  tablename:{guanlianField:"",guanlian2Table:"",guanlian2Field:""}
      let tableGuanlian = this.guanlian;

      for (let i = 0; i < this.curTables.length; i++) {
        let element = this.curTables[i];

        tableRename[element.TableName] = "t" + i;

        if (element.ismore && i != 0) {
          //是否是多个
          tableMore[element.TableName] = element.ismore;
        }
        let fields = [];
        element.Fields.forEach(ff => {
          if (ff.checked == true) {
            fields.push(ff);
          }
        });
        if (fields.length > 0) {
          tableFields[element.TableName] = fields;
          tablenames.push(element.TableName);
        } else {
          this.$message.error("数据表" + element.TableName + "未勾选查询字段");
          return;
        }
        if (i != 0 && tableGuanlian[element.TableName] == null) {
          this.$message.error("数据表" + element.TableName + "未添加关联字段");
          return;
        }
        //检测关联字段是否勾选，关联字段必须勾选
        if (i != 0 && tableGuanlian[element.TableName]) {
          let fields1 = this.curTables[i].Fields;

          let zhuind = 0;
          for (let j = 0; j < this.curTables.length; j++) {
            if (
              this.curTables[j].TableName ==
              tableGuanlian[element.TableName]["guanlian2Table"]
            ) {
              zhuind = j;
              break;
            }
          }
          let fields2 = this.curTables[zhuind].Fields;

          let gfiled1 = tableGuanlian[element.TableName]["guanlianField"];
          let gfiled2 = tableGuanlian[element.TableName]["guanlian2Field"];
          for (let j = 0; j < fields1.length; j++) {
            let e2 = fields1[j];
            if (e2.Name == gfiled1 && e2.checked == false) {
              this.$message.error(
                "关联字段必须勾选:" + element.TableName + "." + gfiled1
              );
              return;
            }
          }
          for (let j = 0; j < fields2.length; j++) {
            let e2 = fields2[j];
            if (e2.Name == gfiled2 && e2.checked == false) {
              this.$message.error(
                "关联字段必须勾选:" +
                  tableGuanlian[element.TableName]["guanlian2Table"] +
                  "." +
                  gfiled2
              );
              return;
            }
          }
        }
      }

      console.log("tables:", tablenames);
      console.log("tables字段:", tableFields);
      console.log("tablemore :", tableMore);
      console.log("table关联:", tableGuanlian);

      //校验 收集结束  下边要生成代码了

      // 查询结构体生成
      let selectziduan = [];
      let selectstruct = "type DBOut struct { \n";
      for (var tname in tableFields) {
        tableFields[tname].forEach(filed => {
          let tmp = "  T" + tname + filed.Name + " ";
          tmp += this.sqltypetogotype(filed.Type) + "  ";
          let tmpziduan = tableRename[tname] + "_" + filed.Name;
          tmp += `"gorm:\\\"column:` + tmpziduan + `\\\"" \n`;
          selectstruct += tmp;

          fieldsRename[tableRename[tname] + "." + filed.Name] = tmpziduan;
        });
      }
      selectstruct += "} \n\n";

      let sqlcontent = "var back []DBOut \n";
      sqlcontent +=
        `err := model.DB.Table(\"` +
        tablenames[0] +
        ` as ` +
        tableRename[tablenames[0]] +
        `\").\n`;
      sqlcontent += `Select(\"`;
      for (var ziduan in fieldsRename) {
        sqlcontent += ziduan + " as " + fieldsRename[ziduan] + ",";
      }
      //去掉最后的逗号
      sqlcontent = sqlcontent.slice(0, sqlcontent.length - 1);
      sqlcontent += `\").\n`;
      for (let i = 1; i < tablenames.length; i++) {
        let tmpgl = tableGuanlian[tablenames[i]];
        let tiaojian =
          tableRename[tablenames[i]] +
          "." +
          tmpgl["guanlianField"] +
          " = " +
          tableRename[tmpgl["guanlian2Table"]] +
          "." +
          tmpgl["guanlian2Field"];
        sqlcontent +=
          `Joins(\"left join ` +
          tablenames[i] +
          ` as ` +
          tableRename[tablenames[i]] +
          ` on ` +
          tiaojian +
          `\").\n`;
      }
      for (let i = 0; i < this.tiaojianSelect.length; i++) {
        sqlcontent += this.tiaojianSelect[i];
      }
      sqlcontent += `Scan(&back).Error\n`;

      sqlcontent += `if err !=nil {
			return nil,err
		  } \n`;

      let structname = this.maptables[tablenames[0]].StructName;

      let funcname = "SelectBy";
      this.tiaojianSelect.forEach((ti, i) => {
        if (this.tiaojianOri[ti] != null) {
          funcname += this.tiaojianOri[ti].OrName;
        }
      });

      let functext = "func (*" + structname + "Dao) " + funcname + "(";
      let canshu = "";
      this.tiaojianSelect.forEach((ti, i) => {
        if (this.tiaojianOri[ti] != null) {
          canshu +=
            this.tiaojianOri[ti].OrName +
            " " +
            this.tiaojianOri[ti].OrType +
            ", ";
        }
      });
      //去掉最后的逗号
      if (canshu.length > 2) {
        canshu = canshu.slice(0, canshu.length - 2);
      }
      let outstructname = structname + funcname + "Back";
      functext += canshu + ") ( []" + outstructname + ",error ){ \n";
      functext += selectstruct;

      functext += sqlcontent;

      //外部结构体
      let OutStruct = "type " + outstructname + " struct { \n";

      // 处理外部struct信息 [{TableName:"", Fields:[{key:"",type:"",tag:""}],links:[]}]
      let tmptablestructs = [];
      tablenames.forEach(tname => {
        let tmpfields = [];
        tableFields[tname].forEach(filed => {
          let key = "T" + tname + filed.Name;
          let type = this.sqltypetogotype(filed.Type);
          let tag = `"json:\\\"` + filed.Name + `\\\""`;
          tmpfields.push({
            key: key,
            type: type,
            tag: tag
          });
        });
        let ttxx = {
          TableName: tname,
          Fields: tmpfields,
          links: []
        };
        tmptablestructs.push(ttxx);
      });
      //根据连表进行联合
      tmptablestructs.forEach(element => {
        tmptablestructs.forEach(e2 => {
          if (
            tableGuanlian[e2.TableName] != null &&
            tableGuanlian[e2.TableName]["guanlian2Table"] == element.TableName
          ) {
            element.links.push(e2);
          }
        });
      });

      // 使用深度优先遍历 合并
      let dfsFields = this.dfs(tmptablestructs[0], tableMore);

      dfsFields.forEach(filed => {
        OutStruct += filed.key + " " + filed.type + " " + filed.tag + "\n";
      });

      OutStruct += "} \n\n";

      let all = "";
      all += OutStruct;
      all += functext;

      //倒腾参数输出

      let havemore = false;
      for (var t in tableMore) {
        if (tableMore[t] == 1) {
          havemore = true;
        }
      }
      let handlestruct = "";

      //如果都是1对1
      if (havemore == false) {
        handlestruct += " out := make([]" + outstructname + ",0) \n";

        handlestruct += `for _, b := range back {\n`;
        handlestruct += "one := " + outstructname + "{}\n";
        for (var tname in tableFields) {
          tableFields[tname].forEach(filed => {
            let tmp = "T" + tname + filed.Name;
            handlestruct += "one." + tmp + " = b." + tmp + "\n";
          });
        }
        handlestruct += "out = append(out,one) \n";
        handlestruct += `} \n`;
      } else {
        handlestruct += "\nout := make([]" + outstructname + ",0) \n";
        handlestruct +=
          "//============== \n //  自己写整合吧 \n //=============== \n";
      }

      all += handlestruct;

      all += "return out, err \n}";

      this.daocode = all;

      //service代码
      let serfunctext = "func (*" + structname + "Service) " + funcname + "(";
      {
        let canshu = "";
        this.tiaojianSelect.forEach((ti, i) => {
          if (this.tiaojianOri[ti] != null) {
            canshu +=
              this.tiaojianOri[ti].OrName +
              " " +
              this.tiaojianOri[ti].OrType +
              ", ";
          }
        });

        //去掉最后的逗号
        if (canshu.length > 2) {
          canshu = canshu.slice(0, canshu.length - 2);
        }
        serfunctext += canshu + ") ( []dao." + outstructname + ",error ){ \n";
        serfunctext += "dao := &dao." + structname + "Dao{} \n";
        serfunctext += "return dao." + funcname + "(  ";
        this.tiaojianSelect.forEach((ti, i) => {
          if (this.tiaojianOri[ti] != null) {
            serfunctext += this.tiaojianOri[ti].OrName + ", ";
          }
        });
        serfunctext = serfunctext.slice(0, serfunctext.length - 2);
        serfunctext += ") \n";
        serfunctext += "}";
      }
      this.servicecode = serfunctext;

      //apicode生成
      let tmpapicode = `
func apinameXXX (c *gin.Context) {

var param struct {\n`;

      this.tiaojianSelect.forEach((ti, i) => {
        if (this.tiaojianOri[ti] != null) {
          tmpapicode +=
            "  P" +
            this.tiaojianOri[ti].OrName +
            " " +
            this.tiaojianOri[ti].OrType +
            ' "json:\\"' +
            this.tiaojianOri[ti].OrName +
            '\\"  ' +
            '  form:\\"' +
            this.tiaojianOri[ti].OrName +
            '\\"" ' +
            "\n";
        }
      });

      tmpapicode +=
        `}
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.` +
        structname +
        `Service{} \n
back, err := ser.` +
        funcname +
        `(  `;
      this.tiaojianSelect.forEach((ti, i) => {
        if (this.tiaojianOri[ti] != null) {
          tmpapicode += "param.P" + this.tiaojianOri[ti].OrName + ", ";
        }
      });
      tmpapicode = tmpapicode.slice(0, tmpapicode.length - 2);
      tmpapicode += `)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
  reply := NewReplyOk()
  reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}
     `;

      this.apicode = tmpapicode;
      this.codedialog = true;

      this.$axios
        .post("/cx/format", {
          code: this.daocode
        })
        .then(res => {
          if (res.data.code == 1) {
            this.daocode = res.data.data;
          }
        });
      this.$axios
        .post("/cx/format", {
          code: this.servicecode
        })
        .then(res => {
          if (res.data.code == 1) {
            this.servicecode = res.data.data;
          }
        });
      this.$axios
        .post("/cx/format", {
          code: this.apicode
        })
        .then(res => {
          if (res.data.code == 1) {
            this.apicode = res.data.data;
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