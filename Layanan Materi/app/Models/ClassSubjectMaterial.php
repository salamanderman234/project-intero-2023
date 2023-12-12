<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class ClassSubjectMaterial extends Model
{
    use HasFactory, SoftDeletes;
    protected $guarded = ["id"];
    protected $primaryKey = 'id';

    public function class_subject()
    {
        return $this->belongsTo(ClassSubject::class,'class_subject_id', 'id');
    }

    public function progres_check()
    {
        return $this->belongsTo(ClassSubjectMaterialCheck::class,'class_subject_material_id', 'id');
    }

    public function student()
    {
        return $this->belongsTo(Student::class,'student_id', 'id');
    }

    public function material_attachments()
    {
        return $this->hasMany(ClassSubjectMaterialAttachment::class,'class_subject_material_id', 'id');
    }

}
