import React from "react";
import { Badge, IconButton } from "@material-ui/core";
import SearchIcon from "@material-ui/icons/Search";
import SearchFloatInput from "./searchFloatInput";
import busEvents from "../../utils/busEvents";

const Search = props => {
  const [showInput, setShowInput] = React.useState(false);
  const [keyword, setKeyword] = React.useState("");

  busEvents.search.setSearchKeyword = setKeyword;
  busEvents.search.getSearchKeyword = () => keyword;

  return (
    <div>
      <IconButton
        edge="start"
        color="primary"
        onClick={() => {
          setShowInput(true);
        }}
      >
        <Badge
          color="secondary"
          variant={keyword.length === 0 ? "standard" : "dot"}
        >
          <SearchIcon />
        </Badge>
      </IconButton>
      <div style={{ display: showInput ? "block" : "none" }}>
        <SearchFloatInput
          closeInput={() => {
            setShowInput(false);
          }}
          keyword={keyword}
          setKeyword={setKeyword}
        />
      </div>
    </div>
  );
};

export default Search;
