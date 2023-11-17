import { styled } from "styled-components";

export const Contain = styled.div`
  margin: 10px;
  border: 1px solid black;
  height: 100px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
`;
export const Title = styled.div`
  background-color: green;
  width: 30%;
  height: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
`;
export const Content = styled.div`
  background-color: red;
  width: 30%;
  height: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
`;
export const CheckBox = styled.div`
  background-color: ${(props) => (props.isChecked ? "black" : "white")};
  width: 10px;
  height: 10px;
  margin-left: 20px;
  border: 1px solid black;
  display: flex;
  justify-content: center;
  align-items: center;
`;
