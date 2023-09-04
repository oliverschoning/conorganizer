"use client";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Button, Card } from "../lib/mui";
import React, { Fragment, useContext, useEffect, useState } from "react";
import { faCoffee } from "@fortawesome/free-solid-svg-icons";
import {
  doc,
  onSnapshot,
  updateDoc,
  setDoc,
  deleteDoc,
  collection,
  serverTimestamp,
  getDocs,
  query,
  where,
  orderBy,
  limit,
} from "firebase/firestore";
import db from "../lib/firebase";
import { AuthContext } from "./auth";
import { Dialog, DialogTitle, IconButton } from "@mui/material";
import AddCircleIcon from "@mui/icons-material/AddCircle";

const AddEvent = () => {
  const [open, setOpen] = React.useState(false);

  const colletionRef = collection(db, "schools");
  const [clicked, setClicked] = useState(false);

  const { currentUser } = useContext(AuthContext);

  const currentUserId = currentUser ? currentUser.uid : null;
  const [schools, setSchools] = useState([]);
  const [loading, setLoading] = useState(false);
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [score, setScore] = useState("");

  useEffect(() => {
    const q = query(
      colletionRef,
      //  where('owner', '==', currentUserId),
      where("title", "==", "School1") // does not need index
      //  where('score', '<=', 100) // needs index  https://firebase.google.com/docs/firestore/query-data/indexing?authuser=1&hl=en
      // orderBy('score', 'asc'), // be aware of limitations: https://firebase.google.com/docs/firestore/query-data/order-limit-data#limitations
      // limit(1)
    );

    setLoading(true);
    const unsub = onSnapshot(colletionRef, (querySnapshot) => {
      const items = [];
      querySnapshot.forEach((doc) => {
        items.push(doc.data());
      });
      setSchools(items);
      setLoading(false);
    });
    return () => {
      unsub();
    };
  }, []);

  // ADD FUNCTION
  const addSchool = async () => {
    const owner = currentUser ? currentUser.uid : "unknown";
    const ownerEmail = currentUser ? currentUser.email : "unknown";

    const newSchool = {
      title,
      desc,
      score: +score,
      owner,
      ownerEmail,
      createdAt: serverTimestamp(),
      lastUpdate: serverTimestamp(),
    };

    try {
      const schoolRef = doc(colletionRef);
      await setDoc(schoolRef, newSchool);
    } catch (error) {
      console.error(error);
    }
  };

  //DELETE FUNCTION
  async function deleteSchool(school) {
    try {
      const schoolRef = doc(colletionRef, school.id);
      await deleteDoc(schoolRef, schoolRef);
    } catch (error) {
      console.error(error);
    }
  }

  // EDIT FUNCTION
  async function editSchool(school) {
    const updatedSchool = {
      score: +score,
      lastUpdate: serverTimestamp(),
    };

    try {
      const schoolRef = doc(colletionRef, school.id);
      updateDoc(schoolRef, updatedSchool);
    } catch (error) {
      console.error(error);
    }
  }


  return (
    <>
      <Dialog
        onClose={() => {
          console.log("close");
        }}
        open={open}
      >
        <DialogTitle>Legg til nytt arangement</DialogTitle>

        <h1>Schools (SNAPSHOT adv.)</h1>
        <div className="inputBox">
          <h3>Add New</h3>
          <h6>Title</h6>
          <input
            className="text-black"
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <h6>Score 0-10</h6>
          <input
            className="text-black"
            type="number"
            value={score}
            onChange={(e) => setScore(e.target.value)}
          />
          <h6>Description</h6>
          <textarea
            className="text-black"
            value={desc}
            onChange={(e) => setDesc(e.target.value)}
          />
          <button onClick={() => addSchool()}>Submit</button>
        </div>

      </Dialog>

      <IconButton
        className="fixed right-0 bottom-0"
        aria-label="add"
        color="error"
        size="large"
        onClick={() => {
          setOpen(true);
        }}
      >
        <AddCircleIcon />
      </IconButton>
    </>
  );
};

export default AddEvent;
