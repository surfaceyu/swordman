// eslint-disable-next-line @typescript-eslint/no-var-requires
import cpy from "cpy";
// 复制配置文件到打包好的目录中
cpy(["./dist/**"], "../backend/public", {parents: true}).then(() => {
    console.log("build. config files copied");
});
