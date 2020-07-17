package askuemssql

import (
  "database/sql"
  "os"
  "log"
  "strconv"
  _ "github.com/denisenkom/go-mssqldb"
  "../storage"
)




var db *sql.DB
//var record string

func ConnectDb() {
    var err error
    mssql_address := os.Getenv("MSSQL_ADDRESS")
    mssql_user := os.Getenv("MSSQL_USER")
    mssql_pass := os.Getenv("MSSQL_PASS")
    db, err = sql.Open("mssql", "server=" + mssql_address + ";user id="+ mssql_user +";password="+ mssql_pass +";")
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

}

func CloseDb() {
  defer db.Close()
}


//возвращает все записи где name содержит заданную подстроку
func ShowTP(arg string) string {
  var s string
  var str string

  var (
    ssName string
    ssGsmPhone string
    msComment string
  )

  if len(arg) != 0 {
    s = "WHERE msUspFile LIKE '%" + arg + "%'"
  }
  rows, err := db.Query("SELECT ssName, ssGsmPhone, msComment FROM askue.dbo.DataSources " + s + " ORDER BY ssName")
  if err != nil {
    log.Fatal(err)
    return "ошибка"
  }
  defer rows.Close()

  for rows.Next() {
    err := rows.Scan(&ssName, &ssGsmPhone, &msComment)
    if err != nil {
      log.Fatal(err)
    }
    str=str+"\n"+ssName +" <b>["+ msComment +"]</b>"

  }

  return str
}

func FindAll(arg string, find string) string {
  var s string
  var str string



  var (
    lId int
    msComment string
    //ssSerialNr string
    //lAbonentId string
    lReserv_3 string

    lReserv_4 string
    dTransKoefI float64

    dVal float64
    dKoefK float64
    tMetrTime string


    msName string
    ssName string
    ssName_tp string
    msComment_tp string
  )


  if len(arg) != 0 {

    if find == storage.SEARCH_BY_ABB {
      s = "WHERE dbo.Abonents.msName LIKE '%" + arg + "%' ORDER BY dbo.Hierar_1_Nodes.ssName"
    }
    //s = " WHERE dbo.Rps.lReserv_3 = "+ arg
    if find == storage.SEARCH_BY_NAME {
      s = "WHERE dbo.Hierar_1_Nodes.ssName LIKE '%" + arg + "%' ORDER BY dbo.Hierar_1_Nodes.ssName"
    }
    if find == storage.SEARCH_BY_NUMBER {
      s = "WHERE dbo.Rps.lReserv_3 = '"+arg+"'"
    }

  }
//  log.Printf("s %s", s)
  rows, err := db.Query("SELECT top 5 dbo.Rps.lId, dbo.Rps.lReserv_3, dbo.Rps.lReserv_4, dbo.Hierar_1_Nodes.ssName, COALESCE(dbo.Abonents.msName,'???') as msName, dbo.Rps.msComment, dbo.Rps.dTransKoefI, dbo.DataSources.ssName as ssName_tp, dbo.DataSources.msComment as msComment_tp FROM dbo.Rps LEFT JOIN dbo.Abonents ON dbo.Abonents.lId=dbo.RPs.lAbonentId LEFT JOIN dbo.Hierar_1_Nodes ON dbo.Hierar_1_Nodes.lRpId=dbo.RPs.lId LEFT JOIN dbo.DataSources ON dbo.DataSources.lId=dbo.RPs.lDsId " + s)
  if err != nil {
    log.Fatal(err)
    return "ошибка"
  }
  defer rows.Close()

  for rows.Next() {


    err := rows.Scan(&lId, &lReserv_3, &lReserv_4, &ssName, &msName, &msComment, &dTransKoefI, &ssName_tp, &msComment_tp)
    if err != nil {
      log.Fatal(err)
      return "ошибка"
    }


    str=str+"\nТП: <b>"+ssName_tp+"</b><i>["+msComment_tp+"]"+"</i>\n<b>"+lReserv_3+"/" + lReserv_4 + " - "+ ssName + "</b> <i>["+ msName + "]</i> " + msComment

//     log.Printf("%s", lAbonentId)
    rows2, err2 := db.Query("SELECT dVal, dKoefK, tMetrTime FROM dbo.RP_"+strconv.Itoa(lId)+"_Params LEFT JOIN dbo.main ON dbo.RP_"+strconv.Itoa(lId)+"_Params.lIdent=dbo.main.lIdent where lParTypeId = 65 or lParTypeId = 12 order by lParTypeId")
    if err2 != nil {
      log.Fatal(err2)
      return "ошибка"
    }
    defer rows2.Close()


     rows2.Next()
      err2 = rows2.Scan(&dVal, &dKoefK, &tMetrTime)
      if err2 != nil {
        log.Fatal(err2)
      }
      if  dVal == 0 {
        str=str+"\n<b>ВЫКЛ</b>"+" <i>["+tMetrTime+"]</i>\n"
      } else{
        str=str+"\n<b>ВКЛ</b>"+" <i>["+tMetrTime+"]</i>\n"
      }

      rows2.Next()

      err2 = rows2.Scan(&dVal, &dKoefK, &tMetrTime)
      tmp := dVal/1000
      str=str+"<b>"+ strconv.FormatFloat(tmp, 'f', 2, 64)+ " кВт*ч </b><i>["+tMetrTime+"]</i>\n"

  }
  return str
}


//
// func main() {
//   ConnectDb()
//   fmt.Println(ShowTP(db,"Овсянка"))
//
//   // var (
//   //   ssName string
//   //   ssGsmPhone string
//   //   msComment string
//   // )
//   // rows, err := db.Query("select ssName, ssGsmPhone, msComment from askue.dbo.DataSources WHERE msUspFile like '%Овсянка%' ORDER BY ssName")
//   // if err != nil {
//   //   log.Fatal(err)
//   // }
//   // for rows.Next() {
//   //   err := rows.Scan(&ssName, &ssGsmPhone, &msComment)
//   //   if err != nil {
//   //     log.Fatal(err)
//   //   }
//   //   fmt.Println(ssName, "\"",msComment,"\"", " : ", ssGsmPhone)
//   // }
//   CloseDb()
// }
