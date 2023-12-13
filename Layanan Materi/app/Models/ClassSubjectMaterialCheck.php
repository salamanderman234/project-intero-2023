<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class ClassSubjectMaterialCheck extends Model
{
    use HasFactory;
    protected $guarded = ["id"];

    public function class_subject_material()
    {
        return $this->belongsTo(ClassSubjectMaterial::class,'class_subject_material_id', 'id');
    }

    public function student()
    {
        return $this->belongsTo(Student::class,'student_id', 'id');
    }
}
