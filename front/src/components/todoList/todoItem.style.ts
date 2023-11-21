import { styled } from "styled-components";

export const Contain = styled.div`
  margin: 10px;
  border: 1px solid black;
  height: 50px;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 100%;
`;
export const Wrap = styled.div`
  display: flex;
  height: 80%;
  align-items: center;
  width: 90%;
`;
export const Title = styled.div`
  background-color: green;
  width: 40%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
`;
export const Content = styled.div`
  background-color: red;
  width: 40%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
`;
export const CheckBox = styled.div`
  background-color: ${(props) => (props.isChecked ? "black" : "white")};
  width: 10px;
  height: 10px;
  border: 1px solid black;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  right: 20px;
`;
