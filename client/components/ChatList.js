"use client";

import React, {useState, useEffect} from "react";
import { useDispatch, useSelector } from "react-redux";
import { getAllAreas } from "@/redux/actions/area";


const ChatList = () => {

    const dispatch = useDispatch();
    const [area, setArea] = useState("");
    const areas = useSelector((state) => state.areaStore.area);
    const translations = {
      "WOMEN_NURSING": "ENFERMERIA MUJERES" ,
      "MEN_NURSING": "ENFERMERIA HOMBRES",
      "AGUARIBAY_NURSING": "ENFERMERIA AGUARIBAY",
      "HOME_DAY_CENTER": "CENTRO DE DIA HOGAR (CDD)",
      "AMBULATORY_DAY_CENTER": "CENTRO DE DIA AMBULATORIO (CDA)",
      "CARE": "ASISTENCIAL",
      "ADMINISTRATION": "ADMINISTRACION",
      "LAUNDRY_AND_WARDROBE": "LAVADERO Y ROPERIA",
      "MAINTENANCE": "MANTENIMIENTO",
      "PHARMACY": "FARMACIA",
      "DOCTORS_OFFICE": "CONSULTORIO MEDICO",
      // Agrega todas las traducciones necesarias
    };


    useEffect(() => {
      dispatch(getAllAreas());
    }, [dispatch]);

    const colors = ["bg-orange-100", "bg-orange-50"];
    return (
      <div className="w-1/3 bg-orange-50 p-4 border-r overflow-y-scroll">
        {areas.length && areas.map((area, index) => (
          <div key={area.id} className={`p-2 text-black hover:bg-gray-200 cursor-pointer ${colors[index % colors.length]}`}>
            {translations[area.name] || area.name} 
          </div>
        ))}
      </div>
    );
  };

  export default ChatList;