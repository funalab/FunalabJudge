import React from "react";
import { PageType } from "../types/PageTypes";
import { Navigate, useLocation, useParams } from "react-router-dom";

type Props = {
  component: React.ReactNode;
  pageType: PageType,
}

export const RouteAuthGuard: React.FC<Props> = (props) => {
  let allowRoute = false;

  const authUserName = localStorage.getItem("authUserName");
  const authJoinedDate = localStorage.getItem("authJoinedDate");
  const authUserExp = localStorage.getItem("authUserExp");

  if ( authUserName && authJoinedDate && authUserExp ) {
    allowRoute = AuthUser({
      pageType: props.pageType,
      authUserName: authUserName,
      authJoinedDate: new Date(authJoinedDate),
      authUserExp: Number(authUserExp),
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
  authJoinedDate: Date,
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
        if (props.authJoinedDate.getFullYear() < new Date().getFullYear()) {
            return true;
        } else {
            if (props.authUserName == userName) {
                return true;  // matched userName
            } else {
                alert("ページへのアクセス権がありません。アクセス権のあるアカウントでログインしてください。");
                return false;  // unmatched userName
            }
        }
    }
    return false  // unknown PageType
}