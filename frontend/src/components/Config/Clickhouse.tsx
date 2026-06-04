import { apiPath } from "../../functions/api"
import { appConfig } from "../../functions/exports"

function Clickhouse() {

  return (
    <div class="card border-primary">
      <div class="card-header">ClickHouse config</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config_clickhouse/'} method="post">
          <table class="table table-borderless"><tbody>
            <tr>
              <td>Enable</td>
              <td>
                <div class="form-check form-switch">
                  {appConfig().ClickhouseEnable
                    ? <input class="form-check-input" type="checkbox" name="enable" checked></input>
                    : <input class="form-check-input" type="checkbox" name="enable"></input>
                  }
                </div>
              </td>
            </tr>
            <tr>
              <td>Address</td>
              <td><input name="addr" type="text" class="form-control" value={appConfig().ClickhouseAddr}></input></td>
            </tr>
            <tr>
              <td>User</td>
              <td><input name="user" type="text" class="form-control" value={appConfig().ClickhouseUser}></input></td>
            </tr>
            <tr>
              <td>Password</td>
              <td><input name="password" type="password" class="form-control" value={appConfig().ClickhousePassword}></input></td>
            </tr>
            <tr>
              <td>Database</td>
              <td><input name="db" type="text" class="form-control" value={appConfig().ClickhouseDB}></input></td>
            </tr>
            <tr>
              <td>Table</td>
              <td><input name="table" type="text" class="form-control" value={appConfig().ClickhouseTable}></input></td>
            </tr>
            <tr>
              <td><button type="submit" class="btn btn-primary">Save</button></td>
              <td></td>
            </tr>
          </tbody></table>
        </form>
      </div>
    </div>
  )
}

export default Clickhouse
