import { useState, createContext, useContext, useEffect } from "react";
import { UserType } from "../types/UserTypes";
import axios from "axios";

axios.defaults.withCredentials = true;

export type AuthUserContextType = {
  user: UserType | null;
  signin: (user:UserType, callback:() => void) => void;
  signout: (callback:() => void) => void;
}
const AuthUserContext = createContext<AuthUserContextType>({} as AuthUserContextType);

export const useAuthUserContext = ():AuthUserContextType => {
  return useContext<AuthUserContextType>(AuthUserContext);
}

type Props = {
  children: React.ReactNode
}

export const AuthUserProvider = (props:Props) => {
  const [user, setUser] = useState<UserType | null>(null);

  const signin = (newUser: UserType, callback: () => void) => {
    setUser(newUser);
    callback();
  }
  const signout = (callback: () => void) => {
    setUser(null);
    callback();
  }
  const value:AuthUserContextType = { user, signin, signout };
  
  return (
    <AuthUserContext.Provider value={value}>
      {props.children}
    </AuthUserContext.Provider>
  );
}