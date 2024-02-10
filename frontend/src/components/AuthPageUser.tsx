import { RoleType, UserType } from "../types/UserTypes";
import { PageType } from "../types/PageTypes";
import { useParams } from "react-router-dom";

type Props = {
  pageType: PageType,
  authUser: UserType,
}

export const AuthPageUser = (props:Props): boolean => {
    // ページの種別ごとにユーザーのアクセス制御を行う関数
    const { userName } = useParams()
    if (props.pageType == PageType.Public) {
        return true;
    } else if (props.pageType == PageType.Private) {
        if (props.authUser.role === RoleType.Admin || props.authUser.role === RoleType.Manager) {  // 書き方綺麗に
            return true;
        } else if (props.authUser.role == RoleType.User) {
            if (props.authUser.userName == userName) {
                return true;
            } else {
                return false;
            }
        } else {
            return false;  // unkown RoleType
        }
    }
    return false  // unkown PageType
}