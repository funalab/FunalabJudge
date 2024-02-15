import React from "react";
import { PageType } from "../types/PageTypes";
import { Navigate, useLocation, useParams } from "react-router-dom";
import { RoleType } from "../types/RoleTypes";

type Props = {
  component: React.ReactNode;
  pageType: PageType,
}

export const RouteAuthGuard: React.FC<Props> = (props) => {
  let allowRoute = false;

  const authUserName = sessionStorage.getItem("authUserName");
  const authUserRole = sessionStorage.getItem("authUserRole");
  const authUserExp = Number(sessionStorage.getItem("authUserExp"));

  if ( authUserName && authUserRole && authUserExp ) {
    allowRoute = AuthUser({
      pageType: props.pageType,
      authUserName: authUserName,
      authUserRole: authUserRole,
      authUserExp: authUserExp,
    });
  } else {
    alert("コンテンツの閲覧にはログインが必要です。");
  }

  if (!allowRoute) {
    return <Navigate to="/login" state={{from:useLocation()}} replace={false} />
  }

  return <>{props.component}</>;

}

type AuthUserProps = {
  pageType: PageType,
  authUserName: string,
  authUserRole: string,
  authUserExp: number
}

export const AuthUser = (props:AuthUserProps): boolean => {
    // ページの種別ごとにユーザーのアクセス制御を行う関数
    const { userName } = useParams()
    if (props.authUserExp < Date.now() / 1000) {
      alert("ログイン保持期限が切れました。再度ログインしてください。")
      return false  // expired token
    }
    if (props.pageType == PageType.Public) {
        return true;
    } else if (props.pageType == PageType.Private) {
        if (props.authUserRole === RoleType.Admin || props.authUserRole === RoleType.Manager) {
            return true;
        } else if (props.authUserRole == RoleType.User) {
            if (props.authUserName == userName) {
                return true;  // matched userName
            } else {
                alert("ページへのアクセス権がありません。アクセス権のあるアカウントでログインしてください。");
                return false;  // unmatched userName
            }
        } else {
            return false;  // unknown RoleType
        }
    }
    return false  // unknown PageType
}